package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/oauth2"
)

// App struct
type App struct {
	ctx context.Context
}

type Release struct {
	Assets []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
	TagName    string `json:"tag_name"`
	PreRelease bool   `json:"prerelease"`
	Published  string `json:"published_at"`
	HtmlURL    string `json:"html_url"`
	Message    string `json:"message"`
}

type ReleaseInfo struct {
	IsLatest bool    `json:"isLatest"`
	Release  Release `json:"release"`
}

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Picture         string `json:"picture"`
	Gender          string `json:"gender"`
	Birthday        string `json:"birthday"`
	Location        string `json:"location"`
	JoinedAt        string `json:"joined_at"`
	AnimeStatistics struct {
		NumItemsWatching    int     `json:"num_items_watching"`
		NumItemsCompleted   int     `json:"num_items_completed"`
		NumItemsOnHold      int     `json:"num_items_on_hold"`
		NumItemsDropped     int     `json:"num_items_dropped"`
		NumItemsPlanToWatch int     `json:"num_items_plan_to_watch"`
		NumItems            int     `json:"num_items"`
		NumDaysWatched      float64 `json:"num_days_watched"`
		NumDaysWatching     float64 `json:"num_days_watching"`
		NumDaysCompleted    float64 `json:"num_days_completed"`
		NumDaysOnHold       float64 `json:"num_days_on_hold"`
		NumDaysDropped      float64 `json:"num_days_dropped"`
		NumDays             float64 `json:"num_days"`
		NumEpisodes         int     `json:"num_episodes"`
		NumTimesRewatched   int     `json:"num_times_rewatched"`
		MeanScore           float64 `json:"mean_score"`
	} `json:"anime_statistics"`
	TimeZone    string `json:"time_zone"`
	IsSupporter bool   `json:"is_supporter"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	if GlobalConfig.MALClientID == "" {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Title:   "Error",
			Message: "Please set your MAL client ID in the config file.",
			Type:    runtime.ErrorDialog,
		})

		os.Exit(1)
	}

	OAuthConfig.ClientID = GlobalConfig.MALClientID

	runtime.EventsOn(ctx, "startup:success", func(data ...interface{}) {

		err := LoadSettings()

		if err != nil {
			runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
				Title:   "Error",
				Message: "Error loading settings.\n\n" + err.Error() + "\n\nSettings will be reset.",
				Type:    runtime.ErrorDialog,
			})

			os.Exit(1)
		}

		jsonString, err := LoadAppDataFile("token.enc", true)

		if err != nil {
			fmt.Println("Error loading token:", err)
		}

		if jsonString != "" {
			AccessToken = &oauth2.Token{}
			err = json.Unmarshal([]byte(jsonString), AccessToken)

			if err != nil {
				fmt.Println("Error unmarshalling token:", err)
			}
		}

		if AccessToken == nil || !AccessToken.Valid() {
			runtime.EventsEmit(ctx, "svelte:goto", "/login")
		}

		err = SetReleaseInfo()

		if err != nil {
			fmt.Println("Error setting release info:", err)
		}

		err = fetchUser()

		if err != nil {
			fmt.Println("Error fetching user:", err)
		}
	})
}

func (a *App) StartAuthProcess() error {
	err := error(nil)
	CodeVerifier, err = GenerateRandomUTFString(128)

	if err != nil {
		fmt.Println("Error generating code verifier:", err)
		return err
	}

	go StartTCPServer(a)

	state, err := GenerateRandomUTFString(32)

	if err != nil {

		return err
	}

	authUrl := OAuthConfig.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("code_challenge", CodeVerifier),
	)

	runtime.BrowserOpenURL(a.ctx, authUrl)

	return nil
}

func (a *App) GetRequest(url string) Response {
	bytes, err := GetRequest(url)

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	return Response{
		Success: true,
		Data:    string(bytes),
	}
}

func (a *App) PatchRequest(url string, json string) Response {
	bytes, err := PatchRequest(url, json)

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	return Response{
		Success: true,
		Data:    string(bytes),
	}
}

func (a *App) DeleteRequest(url string) Response {
	bytes, err := DeleteRequest(url)

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	return Response{
		Success: true,
		Data:    string(bytes),
	}
}

func (a *App) GetVersion() string {
	return Version
}

func (a *App) GetReleaseInfo(prerelease bool) ReleaseInfo {
	return CurrentReleaseInfo
}

func (a *App) InstallUpdate() Response {

	err := DownloadInstaller()

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	err = RunInstaller()

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	os.Exit(0)

	return Response{
		Success: true,
		Data:    "success",
	}
}

func (a *App) SaveSettings(settings string) Response {

	err := json.Unmarshal([]byte(settings), &Settings)

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	err = SaveSettings()

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	return Response{
		Success: true,
		Data:    "success",
	}
}

func (a *App) GetSettings() AppSettings {
	return Settings
}

func (a *App) GetCurrentUser() User {
	return CurrentUser
}

func (a *App) RefreshCurrentUser() Response {
	err := fetchUser()

	if err != nil {
		return Response{
			Success: false,
			Data:    err.Error(),
		}
	}

	return Response{
		Success: true,
		Data:    "success",
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
}

type ReleaseInfo struct {
	IsLatest bool    `json:"isLatest"`
	Release  Release `json:"release"`
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

		jsonString, err := LoadAppDataFile("token.enc")

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

func (a *App) GetRequest(url string) string {

	if AccessToken == nil || !AccessToken.Valid() {
		return "[]"
	}

	fullBearer := "Bearer " + AccessToken.AccessToken

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "error"
	}

	req.Header.Add("Authorization", fullBearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error getting request:", err)
		return "error"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "error"
	}

	return string(body)
}

func (a *App) PatchRequest(url string, json string) string {

	fullBearer := "Bearer " + AccessToken.AccessToken

	formData, err := jsonToFormData(json)

	if err != nil {
		fmt.Println("Error converting json to form data:", err)
		return "error"
	}

	req, err := http.NewRequest("PATCH", url, strings.NewReader(formData.Encode()))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "error"
	}

	req.Header.Add("Authorization", fullBearer)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error getting request:", err)
		return "error"
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "error"
	}

	return string(bodyBytes)
}

func (a *App) GetVersion() string {
	return Version
}

func (a *App) GetReleaseInfo(prerelease bool) ReleaseInfo {
	return CurrentReleaseInfo
}

func (a *App) InstallUpdate() string {

	err := DownloadInstaller()

	if err != nil {
		return err.Error()
	}

	err = RunInstaller()

	if err != nil {
		return err.Error()
	}

	os.Exit(0)

	return "success"
}

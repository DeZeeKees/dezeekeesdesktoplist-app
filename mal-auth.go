package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/oauth2"
)

var (
	redirectURL = "dezeekeesdesktoplist://auth/callback"
)

var OAuthConfig = oauth2.Config{
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://myanimelist.net/v1/oauth2/authorize",
		TokenURL: "https://myanimelist.net/v1/oauth2/token",
	},
	RedirectURL: redirectURL,
}

var AccessToken *oauth2.Token = nil
var CodeVerifier string

var stopServer chan struct{} // Channel to send stop signal
var wg sync.WaitGroup

func StartTCPServer(a *App) {
	listener, err := net.Listen("tcp", ":"+GlobalConfig.IPCPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	defer listener.Close()

	stopServer = make(chan struct{}) // Initialize the stopServer channel

	wg.Add(1)

	go func() {
		defer wg.Done() // Signal the waitgroup when the goroutine is done
		for {
			select {
			case <-stopServer:
				return // Stop the goroutine when receiving a stop signal
			default:
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting: ", err.Error())
					time.Sleep(100 * time.Millisecond)
					continue
				}

				go HandleConnection(conn, a)
			}
		}
	}()

	wg.Wait() // Wait for goroutine to stop
}

func StopTCPServer() {
	close(stopServer) // Send stop signal to goroutine
}

func HandleConnection(conn net.Conn, a *App) {
	defer conn.Close() // Close the connection after handling it

	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	msg := string(buffer[:n])
	fmt.Println("Received message: ", msg)

	token, err := HandleAuthCallback(msg)
	if err != nil {
		fmt.Println("Error handling auth callback:", err.Error())
		return
	}

	AccessToken = token
	StopTCPServer()
	runtime.EventsEmit(a.ctx, "svelte:goto", "/")
}

func SendToRunningInstance(msg string) {
	conn, err := net.Dial("tcp", "localhost:"+GlobalConfig.IPCPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}

	defer conn.Close()

	_, err = conn.Write([]byte(msg))

	if err != nil {
		fmt.Println("Error sending message:", err.Error())
		return
	}

	fmt.Scanln()
}

func HandleAuthCallback(codeString string) (*oauth2.Token, error) {

	parsed, err := url.Parse(codeString)

	if err != nil {
		fmt.Println("Error parsing code:", err)
		return nil, err
	}

	token, err := OAuthConfig.Exchange(
		context.Background(),
		parsed.Query().Get("code"),
		oauth2.SetAuthURLParam("client_id", GlobalConfig.MALClientID),
		oauth2.SetAuthURLParam("code_verifier", CodeVerifier),
		oauth2.SetAuthURLParam("grant_type", "authorization_code"),
	)
	if err != nil {
		fmt.Println("Error exchanging token:", err)
		return nil, err
	}

	//convert token to json
	jsonToken, err := json.Marshal(token)
	if err != nil {
		fmt.Println("Error marshalling token:", err)
		return nil, err
	}

	//save token to appdata
	err = SaveAppDataFile("token.enc", string(jsonToken), true)

	if err != nil {
		fmt.Println("Error saving token:", err)
		return nil, err
	}

	return token, nil
}

func GenerateRandomUTFString(length int) (string, error) {
	// Calculate the required byte length based on the desired string length
	byteLength := (length * 6) / 8

	// Generate a random byte slice of the calculated length
	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	codeVerifier := base64.URLEncoding.EncodeToString(randomBytes)[:length]
	return codeVerifier, nil
}

// jsonToFormData converts a JSON string to url.Values
func jsonToFormData(jsonString string) (url.Values, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}

	formData := url.Values{}
	for key, value := range data {
		formData.Set(key, fmt.Sprint(value))
	}

	return formData, nil
}

package main

import (
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

var appConfig AppConfig

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		SendToRunningInstance(arg)
		return
	}

	appConfig, err := InitConfig()
	if err != nil {
		panic(err)
	}

	if appConfig.MALClientID == "" {
		panic("MALClientID not set in config file")
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "dezeekeesdesktoplist",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Frameless: true,
		Logger:    nil,
		Debug:     options.Debug{},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

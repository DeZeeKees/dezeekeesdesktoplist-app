package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct
type AppConfig struct {
	MALClientID string `toml:"mal_client_id" env:"MAL_CLIENT_ID"`
	IPCPort     string `toml:"ipc_port" env:"IPC_PORT" default:"8499"`
}

var appConfig2 AppConfig

func InitConfig() (AppConfig, error) {

	configFile := GetConfigPath("config")

	//check if congif file exists if not create it
	_, err := os.Stat(configFile)
	if errors.Is(err, os.ErrNotExist) {
		os.Create(configFile)
	}

	// reading config file
	err = cleanenv.ReadConfig(configFile, &appConfig)

	if err != nil {
		fmt.Println(err)
		return appConfig, err
	}

	return appConfig, nil
}

func GetConfigPath(fileName string) string {

	// getting config path
	userConfigDir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println("Cannot find config dir")
	}

	// creating config folder if not exists
	devideConfigDir := filepath.Join(userConfigDir, ".dezeekeesdesktoplist")
	os.Mkdir(devideConfigDir, 0755)

	return filepath.Join(devideConfigDir, fileName+".toml")
}

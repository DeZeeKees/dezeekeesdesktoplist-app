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

func InitConfig() (AppConfig, error) {

	var config AppConfig

	configFile := GetConfigPath("config")

	//check if congif file exists if not create it
	_, err := os.Stat(configFile)
	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(configFile)

		if err != nil {
			fmt.Println(err)
			return config, err
		}

		file.WriteString("mal_client_id = \"\"\nipc_port = \"8499\"")
	}

	// reading config file
	err = cleanenv.ReadConfig(configFile, &config)

	if err != nil {
		fmt.Println(err)
		return config, err
	}

	return config, nil
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

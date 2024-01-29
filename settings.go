package main

import (
	"encoding/json"
)

type AppSettings struct {
	UsePrerelease              bool    `json:"usePrerelease"`
	YourListCardSizeMultiplier float64 `json:"yourListCardSizeMultiplier"`
}

var Settings AppSettings

func LoadSettings() error {

	data, err := LoadAppDataFile("settings.json", false)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), &Settings)

	if err != nil {
		Settings = AppSettings{
			UsePrerelease:              false,
			YourListCardSizeMultiplier: 1,
		}

		err = SaveSettings()

		if err != nil {
			return err
		}
	}

	if Settings.YourListCardSizeMultiplier == 0 {
		Settings.YourListCardSizeMultiplier = 1
	}

	return nil
}

func SaveSettings() error {
	data, err := json.Marshal(Settings)

	if err != nil {
		return err
	}

	err = SaveAppDataFile("settings.json", string(data), false)

	if err != nil {
		return err
	}

	return nil
}

package main

import "encoding/json"

type AppSettings struct {
	UsePrerelease              bool    `json:"usePrerelease"`
	YourListCardSizeMultiplier float64 `json:"yourListCardSizeMultiplier"`
}

var Settings AppSettings

func LoadSettings() error {

	data, err := LoadAppDataFile("settings.enc")

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), &Settings)

	if err != nil {
		return err
	}

	return nil
}

func SaveSettings() error {
	data, err := json.Marshal(Settings)

	if err != nil {
		return err
	}

	err = SaveAppDataFile("settings.enc", string(data))

	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetRequest(url string) ([]byte, error) {
	if AccessToken == nil || !AccessToken.Valid() {
		return nil, fmt.Errorf("invalid token")
	}

	fullBearer := "Bearer " + AccessToken.AccessToken

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fullBearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func PatchRequest(url string, json string) ([]byte, error) {
	if AccessToken == nil || !AccessToken.Valid() {
		return nil, fmt.Errorf("invalid token")
	}

	fullBearer := "Bearer " + AccessToken.AccessToken

	formData, err := jsonToFormData(json)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, strings.NewReader(formData.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fullBearer)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func DeleteRequest(url string) ([]byte, error) {
	if AccessToken == nil || !AccessToken.Valid() {
		return nil, fmt.Errorf("invalid token")
	}

	fullBearer := "Bearer " + AccessToken.AccessToken

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fullBearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func fetchUser() error {
	bytes, err := GetRequest("https://api.myanimelist.net/v2/users/@me")

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &CurrentUser)

	if err != nil {
		return err
	}

	return nil
}

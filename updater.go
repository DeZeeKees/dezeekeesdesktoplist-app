package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func getLatestRelease(preRelease bool) (Release, error) {

	var release Release
	var err error

	if preRelease {
		release, err = getLatestPreRelease()
	} else {
		release, err = getLatestNonPreRelease()
	}

	if err != nil {
		return Release{}, err
	}

	return release, nil

}

func getLatestPreRelease() (Release, error) {
	resp, err := http.Get("https://api.github.com/repos/DeZeeKees/dezeekeesdesktoplist-app/releases")
	if err != nil {
		return Release{}, err
	}
	defer resp.Body.Close()

	var releases []Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return Release{}, err
	}

	for _, release := range releases {
		if release.PreRelease {
			return release, nil
		}
	}

	return Release{}, nil
}

func getLatestNonPreRelease() (Release, error) {
	resp, err := http.Get("https://api.github.com/repos/DeZeeKees/dezeekeesdesktoplist-app/releases/latest")
	if err != nil {
		return Release{}, err
	}
	defer resp.Body.Close()

	var releases Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return Release{}, err
	}

	return Release{}, nil
}

func SetReleaseInfo() error {
	release, err := getLatestRelease(true)

	if err != nil {
		fmt.Println("Error getting latest release:", err)
		return err
	}

	CurrentReleaseInfo = ReleaseInfo{
		IsLatest: release.TagName == Version,
		Release:  release,
	}

	return nil
}

func DownloadInstaller() error {
	resp, err := http.Get("https://api.github.com/repos/DeZeeKees/dezeekeesdesktoplist-installer/releases/latest")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var releases Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return err
	}

	for _, asset := range releases.Assets {
		if asset.Name == "dezeekeesdesktoplist-installer.exe" {
			err := DownloadFile(asset.BrowserDownloadURL, asset.Name)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func DownloadFile(url string, assetName string) error {

	executablePath, err := os.Executable()

	if err != nil {
		return err
	}

	executableDir := filepath.Dir(executablePath)

	out, err := os.Create(executableDir + "/" + assetName)

	if err != nil {
		return err
	}

	defer out.Close()

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}

func RunInstaller() error {

	executablePath, err := os.Executable()

	if err != nil {
		return err
	}

	executableDir := filepath.Dir(executablePath)

	verb := "runas"
	exe := executableDir + "/dezeekeesdesktoplist-installer.exe"
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[:1], " ") + " --update --prerelease"

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err = windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

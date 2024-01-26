package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

const EncryptSecret = "2465df79a950b7b5ef83ee16a93e368a"

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func GetAppDataPath() (string, error) {
	userConfigDir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println("Cannot find config dir")
		return "", err
	}

	appdataPath := filepath.Join(userConfigDir, ".dezeekeesdesktoplist")
	os.Mkdir(appdataPath, 0755)

	return appdataPath, nil
}

func GetAppDataFilePath(filename string) (string, error) {
	appdataPath, err := GetAppDataPath()

	if err != nil {
		fmt.Println("Cannot find appdata path")
		return "", err
	}

	return filepath.Join(appdataPath, filename), nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return data, err
	}
	return data, nil
}

func EncryptString(str string, EncryptSecret string) (string, error) {
	block, err := aes.NewCipher([]byte(EncryptSecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(str)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func DecryptString(str string, EncryptSecret string) (string, error) {
	block, err := aes.NewCipher([]byte(EncryptSecret))
	if err != nil {
		return "", err
	}

	cipherText, err := Decode(str)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func SaveAppDataFile(filename string, data string, encrypt bool) error {
	filepath, err := GetAppDataFilePath(filename)

	if err != nil {
		fmt.Println("Cannot find file")
		return err
	}

	if encrypt {
		//encrypt file
		data, err = EncryptString(data, EncryptSecret)

		if err != nil {
			fmt.Println("Cannot encrypt file")
			return err
		}
	}

	//write file
	_, err = os.Stat(filepath)

	var file *os.File

	if os.IsNotExist(err) {
		//create file
		file, err = os.Create(filepath)
	} else {
		//open file
		file, err = os.OpenFile(filepath, os.O_RDWR, 0644)
	}

	if err != nil {
		fmt.Println("Cannot open file")
		return err
	}

	defer file.Close()

	//clear file
	err = file.Truncate(0)

	if err != nil {
		fmt.Println("Cannot clear file")
		return err
	}

	_, err = file.WriteString(data)

	if err != nil {
		fmt.Println("Cannot write to file")
		return err
	}

	return nil
}

func LoadAppDataFile(filename string, encrypt bool) (string, error) {

	filepath, err := GetAppDataFilePath(filename)

	if err != nil {
		fmt.Println("Cannot get filepath")
		return "", err
	}

	fmt.Println("filepath - loadappdatafile: " + filepath)

	//open file
	_, err = os.Stat(filepath)

	if os.IsNotExist(err) {
		return "", nil
	}

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Cannot open file")
		return "", err
	}

	defer file.Close()

	//read file
	buf := make([]byte, 4096)
	n, err := file.Read(buf)

	if err != nil {
		fmt.Println("Cannot read file")
		return "", err
	}

	fileData := string(buf[:n])

	if encrypt {
		//decrypt file
		fileData, err = DecryptString(fileData, EncryptSecret)

		if err != nil {
			fmt.Println("Cannot decrypt file")
			return "", err
		}
	}

	return fileData, nil
}

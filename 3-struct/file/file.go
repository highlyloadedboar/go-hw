package file

import (
	"errors"
	"os"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return "", errors.New("Error while opening file! " + err.Error())
	}

	defer file.Close()

	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", errors.New("Error while reading fil! " + err.Error())
	}

	return string(content), nil

}

func WriteJsonToFile(fileName string, content []byte) (string, error) {

	err := os.WriteFile(fileName, content, 0644)

	if err != nil {
		return "", err
	}

	return fileName, nil

}

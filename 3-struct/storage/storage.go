package storage

import (
	"demo/3-struct/bins"
	"demo/3-struct/file"
	"encoding/json"
	"os"
)

func SaveBin(fileName string, bin *bins.Bin) (string, error) {
	content, err := json.Marshal(bin)

	file.WriteJsonToFile(fileName, content)

	if err != nil {
		return "", err
	}

	return fileName, nil
}

func GetAllBins(fileName string) (string, error) {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

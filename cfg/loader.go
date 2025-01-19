package cfg

import (
	"encoding/json"
	"io"
	"os"
)

func LoadConfig() (*ApiConfig, error) {
	permissions := 0444 // -r-r-r
	configFile, err := os.OpenFile("cfg/app.json", os.O_RDONLY, os.FileMode(permissions))

	if err != nil {
		return nil, err
	}

	defer configFile.Close()

	data, err := io.ReadAll(configFile)

	if err != nil {
		return nil, err
	}

	var config ApiConfig

	err = json.Unmarshal(data, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

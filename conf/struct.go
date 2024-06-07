package conf

import (
	"encoding/json"
	"os"
)

type Config struct {
	ContestDir string `json:"ContestDir"`
}

func ReadConfig(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := Config{}
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

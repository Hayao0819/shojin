package conf

import (
	"errors"
	"path"
)

var contestDir = ""

func GetContestDir() (string, error) {
	if contestDir == "" {
		return "", errors.New("not initialized")
	}
	return contestDir, nil
}

func Initialize() error {
	baseDir, err := FindBaseDir()
	if err != nil {
		return err
	}

	config, err := ReadConfig(path.Join(*baseDir, configName))
	if err != nil {
		return err
	}

	if path.IsAbs(config.ContestDir) {
		contestDir = config.ContestDir
	} else {
		contestDir = path.Join(*baseDir, config.ContestDir)
	}
	return nil
}

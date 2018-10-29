package database

import (
	"encoding/json"
	"os"
)

const defaultConfigPath = "config/dbConfig.json"

type DbConfig struct {
	DbVendor string `json:"dbVendor"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitWithDefaultConfigs() (*DbConfig, error) {
	decoder, err := getDecoder(defaultConfigPath)
	if err != nil {
		return nil, err
	}

	return createDbConfig(decoder)
}

func InitWithConfigPath(path string) (*DbConfig, error) {
	decoder, err := getDecoder(path)
	if err != nil {
		return nil, err
	}

	return createDbConfig(decoder)
}

func getDecoder(path string) (*json.Decoder, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(jsonFile)

	return decoder, nil
}

func createDbConfig(decoder *json.Decoder) (*DbConfig, error) {
	var dbConfig *DbConfig
	if err := decoder.Decode(&dbConfig); err != nil {
		return nil, err
	}

	return dbConfig, nil
}

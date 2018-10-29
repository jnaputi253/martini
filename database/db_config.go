package database

import (
	"encoding/json"
	"fmt"
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

	var dbConfig *DbConfig
	if err = decoder.Decode(&dbConfig); err != nil {
		return nil, err
	}

	fmt.Printf("DB Vendor: %s\nUsername: %s\nPassword: %s\n",
		dbConfig.DbVendor,
		dbConfig.Username,
		dbConfig.Password)

	return dbConfig, nil
}

func getDecoder(path string) (*json.Decoder, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(jsonFile)

	return decoder, nil
}

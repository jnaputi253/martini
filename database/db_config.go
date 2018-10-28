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
	jsonParser, err := getJsonParser(defaultConfigPath)

	if err != nil {
		return nil, err
	}

	var dbConfig *DbConfig
	if err = jsonParser.Decode(&dbConfig); err != nil {
		return nil, err
	}

	fmt.Printf("DB Vendor: %s\nUsername: %s\nPassword: %s\n",
		dbConfig.DbVendor,
		dbConfig.Username,
		dbConfig.Password)

	return dbConfig, nil
}

func getJsonParser(path string) (*json.Decoder, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(jsonFile)

	return jsonParser, nil
}

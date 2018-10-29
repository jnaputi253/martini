package database

import (
	"encoding/json"
	"os"
)

const defaultConfigPath = "config/dbConfig.json"

// DbConfig contains the information used to connect to a database.
type DbConfig struct {
	DbVendor *string `json:"dbVendor"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// InitWithDefaultConfigs creates an instance of the DbConfig struct by reading the database vendor
// and connection information from a configuration file.
//
// By default, the name of the file is "dbConfig.json" and it is located in the "config/" directory in the root of this project.
// If the configuration file has a different name, or is in another directory, then use the InitWithConfigPath function
// to specify the path to the file.
func InitWithDefaultConfigs() (*DbConfig, error) {
	decoder, err := getDecoder(defaultConfigPath)
	if err != nil {
		return nil, err
	}

	return createDbConfig(decoder)
}

// InitWithConfigPath creates an instance of the DbConfig struct by reading the databse vendor
// and connection information from the file specified in the path parameter.
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

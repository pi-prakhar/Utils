package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/pi-prakhar/utils/logger"
)

var Logger = log.New(log.DEBUG, "Utils")

// LoadEnv loads environment variables from a .env file
func LoadEnv() error {
	err := CheckEnvFile()
	if err != nil {
		Logger.Debug("Error: env file not found")
		return err
	}
	err = godotenv.Load()
	if err != nil {
		Logger.Debug("Error: Fail to load env")
		return err
	}
	return nil
}

// CheckEnvFile checks if a .env file is present in the current directory
func CheckEnvFile() error {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		Logger.Debug("Error: env file not found")
		return err
	}
	return nil
}

func LoadConfig() (map[string]interface{}, error) {

	// Replace with the path to your config file
	configFile := "config/config.json"

	// Read the JSON file
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		Logger.Debug("Error : Failed to read Config file at path config/config.json")
		return nil, err
	}

	// Define a map to store the JSON data
	data := make(map[string]interface{})

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		Logger.Debug(fmt.Sprintf("Failed to decode Config file at path : %s", configFile))
		return nil, err
	}

	return data, nil
}

func GetValueFromConf(key string) (string, error) {
	config, err := LoadConfig()
	if err != nil {
		Logger.Debug("Error : Failed to Load Config File")
		return "", err
	}

	value, found := config[key]
	if !found {
		Logger.Debug(fmt.Sprintf("Key '%s' not found in config file", key))
		return "", fmt.Errorf("key '%s' not found in config file", key)
	}

	// Perform type assertion only if the key is found
	strValue, ok := value.(string)
	if !ok {
		Logger.Debug(fmt.Sprintf("Error : Value for key '%s' is not a string", key))
		return "", fmt.Errorf("value for key '%s' is not a string", key)
	}

	return strValue, nil
}

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl				string	`json:"db_url"`
	CurrentUserName 	string	`json:"current_user_name"`
}

func (c Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	err := writeConfig(c)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	var cfg Config
	configFilePath, err := getConfigFilePath()
	// fmt.Print(configFilePath)
	if err != nil {
		return cfg, err
	}

	fileBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return cfg, err
	}

	// fmt.Printf("Current User Name: %s\nDB URL: %s\n", cfg.CurrentUserName, cfg.DbUrl)

	return cfg, nil


} 
	
func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	} 
	return homedir + "/" + configFileName, nil
}

func writeConfig(cfg Config) error {
	fileBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// 3. Write the bytes to a file
	// 0644 sets standard read/write permissions for the file owner
	configFilePath, err := getConfigFilePath()
	// fmt.Print(configFilePath)
	if err != nil {
		return fmt.Errorf("Error getting config file path: %v\n", err)
	}
	err = os.WriteFile(configFilePath, fileBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("Successfully saved configuration to %s\n", configFilePath)
	return nil
}
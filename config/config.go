package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	configFolder = os.Getenv("HOME") + "/.config/devnotes"
	configPath   = configFolder + "/config.yaml"
	dbPath       = configFolder + "/devnotes.db"
)

const (
	defaultDailySummaryFile = "%Y/%m/%d.md"
)

// ----------------------------

type Config struct {
	DailySummaryFile string `yaml:"daily_summary_file"`
	DatabasePath     string `yaml:"database_path"`
}

func defaultConfig() *Config {
	return &Config{
		DailySummaryFile: defaultDailySummaryFile,
		DatabasePath:     dbPath,
	}
}

func LoadConfig() (*Config, error) {
	c := defaultConfig()

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(configFolder, os.ModePerm)

			err = c.Save()
			if err != nil {
				return nil, fmt.Errorf("Config file not found, failed to create a new one at %s", configPath)
			}

			return nil, fmt.Errorf("Config file not found, created a new one at %s", configPath)
		}
		return nil, err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) Save() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, data, 0644)
}

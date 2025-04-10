package settings

import (
	"os"

	"github.com/kkrypt0nn/overscry/models"
	"gopkg.in/yaml.v3"
)

type Settings struct {
	Version     string      `yaml:"version"`
	Author      string      `yaml:"author"`
	Description string      `yaml:"description"`
	Node        models.Node `yaml:"node"`
}

func LoadSettings(fileName string) (*Settings, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var s Settings
	err = yaml.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func LoadSettingsFromString(yml string) (*Settings, error) {
	var s Settings
	err := yaml.Unmarshal([]byte(yml), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

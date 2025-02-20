package config

import (
    "os"
    "gopkg.in/yaml.v3"
)

type ServerConfig struct {
    Port  int  `yaml:"port"`
    Debug bool `yaml:"debug"`
}

type Route struct {
    Name   string `yaml:"name"`
    Path   string `yaml:"path"`
    Target string `yaml:"target"`
}

type Config struct {
    Server ServerConfig `yaml:"server"`
    Routes []Route      `yaml:"routes"`
}

func LoadConfig(path string) (*Config, error) {
    config := &Config{}
    file, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(file, config)
    return config, err
}
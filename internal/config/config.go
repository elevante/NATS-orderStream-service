package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type NatsStreamingConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	ClusterID string `yaml:"сlusterID"`
	ClientID  string `yaml:"clientID"`
}

type AppConfig struct {
	Database StorageConfig       `yaml:"database"`
	Stan     NatsStreamingConfig `yaml:"stan"`
}

func GetConfig() (AppConfig, error) {
	var config AppConfig

	file, err := os.ReadFile("/home/user/Desktop/WBTECH/NATS-OrderStream-Service/config.yaml")

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Println(err)
	}
	return config, err
}

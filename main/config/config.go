package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var Cfg *BasketConfig

type BasketConfig struct {
	Server struct {
		Port   string `yaml:"port"`
		Host   string `yaml:"host"`
		Secret string `yaml:"secret"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"name"`
		Host     string `yaml:"host"`
	} `yaml:"database"`

	Route struct {
		Basket string `yaml:"basket"`
		Login  string `yaml:"login"`
		Signup string `yaml:"signup"`
	} `yaml:"route"`
}

func ReadConfig() (*BasketConfig, error) {
	Cfg = &BasketConfig{}
	err := cleanenv.ReadConfig("config.yaml", Cfg)
	if err != nil {
		log.Println("Error reading configuration:", err)
		return &BasketConfig{}, err
	}

	return Cfg, nil
}

package config

import (
    "os"

    "gopkg.in/yaml.v2"
)

const configPath = "config/config.yml"

type Config struct {
    Server struct {
        Port string `yaml:"port"`
    } `yaml:"server"`

    Database struct {
        DSN string `yaml:"dsn"`
    } `yaml:"database"`

	Redis struct {
		Addr     string `yaml:"addr"`     
		Password string `yaml:"password"` 
		DB       int    `yaml:"db"`    
	} `yaml:"redis"`
}
var Cfg Config

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)

	if err != nil {
		panic(err)
	}
}
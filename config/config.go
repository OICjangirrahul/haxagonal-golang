package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, defaulting to system env variables")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %v", err))
	}

	configText := os.ExpandEnv(string(data))

	err = yaml.Unmarshal([]byte(configText), &Cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse config: %v", err))
	}

	Cfg.Server.Port = getEnv("PORT", Cfg.Server.Port)
	Cfg.Database.DSN = getEnv("DSN", Cfg.Database.DSN)
	Cfg.Redis.Addr = getEnv("REDIS_ADDR", Cfg.Redis.Addr)
	Cfg.Redis.Password = getEnv("REDIS_PASSWORD", Cfg.Redis.Password)
	if db, exists := os.LookupEnv("REDIS_DB"); exists {
		Cfg.Redis.DB, _ = strconv.Atoi(db)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

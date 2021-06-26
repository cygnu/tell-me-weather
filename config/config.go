package config

import (
	"gopkg.in/go-ini/ini.v1"
	"log"
	"os"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	Port      int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("openweather").Key("api_key").String(),
		ApiSecret: cfg.Section("openweather").Key("api_secret").String(),
		Port:      cfg.Section("web").Key("port").MustInt(),
	}
}

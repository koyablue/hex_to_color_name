package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	BaseUrl string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		BaseUrl: cfg.Section("api").Key("base_url").String(),
	}
}

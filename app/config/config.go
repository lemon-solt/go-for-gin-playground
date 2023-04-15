package config

import (
	"app/utils"
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port    string
	LogFile string
}

var Config ConfigList

func init() {
	// LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("settings.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").MustString("8082"),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}
}

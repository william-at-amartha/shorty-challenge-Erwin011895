package config

import (
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

type Config struct {
	Environment string `ini:"environment"`
	HttpPort string `ini:"httpport"`
}

var appConfig *Config

// Init private instance config
func Init() {
	p, _ := filepath.Abs("./config/app.ini")
	paths := strings.SplitAfter(p, "shorty-challenge")
	
	p = paths[0] + "/config/app.ini"

	appConfig = &Config{}
	iniConfig, err := ini.Load(p)
	if err != nil {
		log.Fatalf("[Init] failed to read config, %+v\n", err)
	}
	err = iniConfig.MapTo(appConfig)
	if err != nil {
		log.Fatalf("[Init] failed to mapping config, %+v\n", err)
	}
}

// Get private instance config. Singleton pattern
func Get() *Config {
	return appConfig
}

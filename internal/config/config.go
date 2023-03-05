package config

import (
	"go-gin-learning/pkg/util"
	"log"

	"github.com/go-ini/ini"
)

type AppCfg struct {
	RootPath string

	LogPath   string
	LogFormat string
	LogSuffix string

	RunMode string
	RunPort       int
	ReadTimeout   int
	WriteTimeout  int
	MaxHeaderBytes int
}

type DBCfg struct {
	Type string
	Path string
	File string

	Name     string
	Password string
}

type CacheCfg struct {
	Type     string
	Password string
}

var App = &AppCfg{}
var DB = &DBCfg{}
var Cache = &CacheCfg{}

func Init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Parsing the config failed, %v", err)
	}
	util.IniToStruct(cfg, "app", App)
	util.IniToStruct(cfg, "db", DB)
	util.IniToStruct(cfg, "cache", Cache)
}

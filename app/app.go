package app

import (
	"log"

	"gopkg.in/ini.v1"
)

// App configuration for app
type App struct {
	PageSize       int
	JwtSecret      string
	PrefixURL      string
	RntimeRootPath string
}

// AppSetting => app setting
var AppSetting = &App{}

// DatabaseSetting => database setting
var DatabaseSetting = &Database{}

// Setup initialize configration
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("")
	}

	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
	InitDB()
}

var cfg *ini.File

// mapTo => map ini file field to struct
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

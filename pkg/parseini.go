package pkg

import (
	"flag"
	"github.com/go-ini/ini"
	"log"
)

type App struct {
	Sql 			string
	Filename		string
	OverwriteFile	bool
	FilesPath 		string
}

var AppSetting = &App{}


type Database struct {
	Type			string
	SqliteFile 		string
	User			string
	Password		string
	Host 			string
	DbName 			string
	Port 			int
	SslMode 		string
}

var DatabaseSetting = &Database{}


var cfg *ini.File

func SetupINI() {
	var err error

	file := confFile()

	cfg, err = ini.Load(file)
	if err != nil {
		log.Fatalf("Setup, fail to parse 'db2csv.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)


}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

func confFile() string {

	confPtr := flag.String("conf", "", "set private app.ini file")
	flag.Parse()

	if *confPtr != "" {
		return *confPtr
	} else {
		return "db2csv.ini"
	}
}

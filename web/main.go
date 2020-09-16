package main

import (
	"fmt"
	"log"

	"github.com/im-jinsu/yepanmap/shared/loadconf"
	"github.com/im-jinsu/yepanmap/web/server"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	loadconf.SetRootConfig()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/log/out.log", loadconf.WASROOTDIR),
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     5, //days
	})
	loadconf.SetConfigController()
	log.Println("Start Yepanmap")
	server.Run()
}

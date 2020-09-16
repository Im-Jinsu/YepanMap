package loadconf

import (
	"log"
	"os"
)

// ConfigBlock : Setting File Block
type ConfigBlock struct {
	MONGODB struct {
		IPPort   string
		DBUser   string
		Password string
	}
}

// WASROOTDIR is Root directory
var WASROOTDIR string

// SetRootConfig load WASROOTDIR
func SetRootConfig() {
	if WASROOTDIR = os.Getenv("WASROOTDIR"); WASROOTDIR == "" {
		log.Println("[ERR] Please check your environment variables. [WASROOTDIR]")
		os.Exit(1)
	}
}

//ConfigInfo : Setting
var ConfigInfo = ConfigBlock{}

// SetConfigController : Config Load
func SetConfigController() {
	if ConfigInfo.MONGODB.IPPort = os.Getenv("MGOIP"); ConfigInfo.MONGODB.IPPort == "" {
		log.Println("[ERR] Please check your environment variables. [MGOIP]")
		os.Exit(1)
	}
	if ConfigInfo.MONGODB.DBUser = os.Getenv("MGOID"); ConfigInfo.MONGODB.DBUser == "" {
		log.Println("[ERR] Please check your environment variables. [MGOID]")
		os.Exit(1)
	}
	if ConfigInfo.MONGODB.Password = os.Getenv("MGOPD"); ConfigInfo.MONGODB.Password == "" {
		log.Println("[ERR] Please check your environment variables. [MGOPD]")
		os.Exit(1)
	}
}

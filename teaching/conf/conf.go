package conf

import (
	"encoding/json"
	"io/ioutil"

	"teach.me/teaching/tlog"
)

type Config struct {
	Server    string `json:"server"`
	Port      int    `json:"port"`
	MgoServer string `json:"mgoServer"`
	MgoPort   int    `json:"mgoPort"`
	TopLimit  int    `json:"topLimit"`
	PageSize  int    `json:"pageSize"`
	LogFile   string `json:"logFile"`
}

var Gconfig Config

func SetConfig() {
	tlog.Info("set config from config.json")
	bys, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		tlog.Fatal(err)
	}
	err = json.Unmarshal(bys, &Gconfig)
	if err != nil {
		tlog.Fatal(err)
	}
}

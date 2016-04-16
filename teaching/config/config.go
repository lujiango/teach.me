package config

import (
	"encoding/json"
	"io/ioutil"

	"teach.me/teaching/tlog"
)

type Config struct {
	Server      string `json:"server"`
	Port        int    `json:"port"`
	LogFile     string `json:"logFile"`
	ConfFile    string `json:"confFile"`
	MgoServer   string `json:"mgoServer"`
	MgoPort     int    `json:"mgoPort"`
	TopLimit    int    `json:"topLimit"`
	PageSize    int    `json:"pageSize"`
	QiNiuDomain string `json:"qiNiuDomain"`
	Bucket      string `json:"bucket"`
	Ak          string `json:"ak"`
	Sk          string `json:"sk"`
}

var Gconfig Config

func SetConfig() {
	tlog.Info("set config from config.json")
	bys, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		tlog.Fatal(err)
	}
	err = json.Unmarshal(bys, &Gconfig)
	if err != nil {
		tlog.Fatal(err)
	}
}

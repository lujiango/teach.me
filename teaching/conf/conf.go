package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Server    string `json:"server"`
	Port      int    `json:"port"`
	MgoServer string `json:"mgoServer"`
	MgoPort   int    `json:"mgoPort"`
	TopLimit  int    `json:"topLimit"`
	PageSize  int    `json:"pageSize"`
}

var gConfig Config

func SetConfig() {
	bys, err := ioutil.ReadFile("teaching.conf")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bys, &gConfig)
	if err := nil{
		log.Fatalln(err)
	}
}

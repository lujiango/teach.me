package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"os"
)

// struct to byte
func Struct2Byte(any interface{}) ([]byte, error) {
	s, err := Struct2Json(any)
	if err != nil {
		return []byte(""), err
	}
	return []byte(s), nil

}

// byte to struct
func Byte2Struct(data []byte, any interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(any)
}

//json to map
func Json2Map(jsonString string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &dat)
	return dat, err
}

//json to struct
func Json2Struct(jsonString string, any *interface{}) error {
	err := json.Unmarshal([]byte(jsonString), any)
	return err
}

//struct to json str
func Struct2Json(any interface{}) (string, error) {
	b, err := json.Marshal(any)
	return string(b), err
}

// struct to map
func Struct2Map(any interface{}) (map[string]interface{}, error) {
	var dat map[string]interface{}
	b, err := Struct2Byte(any)
	if err != nil {
		return nil, err
	}
	log.Println("b >>> ", string(b))
	err = json.Unmarshal(b, &dat)
	log.Println("dat >>> ", dat)
	return dat, err
}

//map to json str
func Map2Json(dat map[string]interface{}) string {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)
	return ""
}

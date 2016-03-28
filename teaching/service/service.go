package service

import (
	"encoding/json"
	"log"

	"teach.me/teaching/mongo"
)

func GetCoursesByLocation(location string) string {
	courses := mongo.QueryAll(mongo.COURSE_COLL, location)
	log.Println(courses)
	bytes, err := json.Marshal(courses)
	if err != nil {
		log.Println(err)
		return ""
	}
	var c = string(bytes)
	var res string
	res = "{\"top_courses\":" + c + ",\"data\":" + c + ",\"timestamp\":1459051200}"
	return res
}

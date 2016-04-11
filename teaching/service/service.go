package service

import (
	"encoding/json"
	"log"
	"strconv"

	"gopkg.in/mgo.v2/bson"
	"teach.me/teaching/mongo"
	"teach.me/teaching/utils"
)

func GetCoursesByLocation(location string, timestamp int64) string {
	var retString string = ""
	var topString string = ""
	var itemString string = ""
	var dataString string = ""
	var ts int = 1460287343
	if timestamp == 0 {
		//{top_courses:[],items:[],data:[],timestamp:14xxxxxxx}
		//step 1: get top_courses
		var tops []mongo.Course
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "istop": 1}).All(&tops)
		log.Println("tops >>> ", tops)
		for _, v := range tops {
			log.Println("v >>> ", v)
			dat, err := utils.Struct2Map(v)
			log.Println("dat >>> ", dat)
			if err != nil {
				log.Println(err)
				return ""
			}

			var teacher mongo.Teacher
			mongo.GetCollection(mongo.TEACHER_COLL).Find(bson.M{"tid": v.TeacherId}).One(&teacher)
			log.Println("teacherId >>> ", v.TeacherId)
			log.Println("teacher >>> ", teacher)
			// remove teacherid
			delete(dat, "TeacherId")
			dat["teacher"] = teacher
			b, _ := json.Marshal(dat)
			topString = topString + string(b) + ","
		}
		// step 2: get items
		var items []mongo.Item
		mongo.GetCollection(mongo.ITEM_COLL).Find(bson.M{}).All(&items)
		for _, i := range items {
			bi, _ := json.Marshal(i)
			itemString = itemString + string(bi) + ","
		}

		// step 3: get data
		var data []mongo.Course
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "isTop": 0}).All(&data)
		for _, d := range data {
			log.Println("d >>> ", d)
			dat, err := utils.Struct2Map(d)
			log.Println("dat >>> ", dat)
			if err != nil {
				log.Println(err)
				return ""
			}
			// remove teacherid
			delete(dat, "teacherId")
			var teacher mongo.Teacher
			mongo.GetCollection(mongo.TEACHER_COLL).Find(bson.M{"tid": d.TeacherId}).One(&teacher)

			log.Println(teacher)
			dat["teacher"] = teacher
			bd, _ := json.Marshal(dat)
			dataString = dataString + string(bd) + ","
		}
	} else {
		//{data:[],timestamp:14xxxxxxx}
		//		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location}).All(&courses)
	}
	retString = "{\"top_courses\":[" + topString + "],\"data\":[" + dataString + "],\"timestamp\":" + strconv.Itoa(ts) + "}"
	log.Println("retString >>> ", retString)
	return retString
}

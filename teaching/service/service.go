package service

import (
	"encoding/json"
	"log"
	"strconv"

	"gopkg.in/mgo.v2/bson"
	"teach.me/teaching/mongo"
)

func GetCoursesByLocation(location string, timestamp int64) string {
	var retString string = ""
	var topString string = ""
	var itemString string = ""
	var dataString string = ""
	var ts int = 0
	log.Println("timestamp", timestamp)
	if timestamp == 0 {
		log.Println("entering if...")
		//{top_courses:[],items:[],data:[],timestamp:14xxxxxxx}
		//step 1: get top_courses
		var tops []bson.M
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "isTop": 1}).Limit(mongo.TOP_LIMIT).All(&tops)
		log.Println("tops.size >>> ", len(tops))
		for _, t := range tops {
			var teacher bson.M
			mongo.GetCollection(mongo.TEACHER_COLL).Find(bson.M{"tid": t["teacherId"]}).One(&teacher)
			// remove teacherId and _id
			delete(t, "teacherId")
			delete(t, "_id")
			delete(teacher, "_id")
			t["teacher"] = teacher
			b, _ := json.Marshal(t)
			topString = topString + string(b) + ","
		}
		topString = topString[:len(topString)-1]
		// step 2: get items
		var items []bson.M
		mongo.GetCollection(mongo.ITEM_COLL).Find(nil).All(&items)
		log.Println("items.size >>> ", len(items))
		for _, i := range items {
			//remove _id
			delete(i, "_id")
			bi, _ := json.Marshal(i)
			itemString = itemString + string(bi) + ","
		}
		itemString = itemString[:len(itemString)-1]
		// step 3: get data
		var data []bson.M
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "isTop": 0}).Limit(mongo.PAGE_SIZE).All(&data)
		log.Println("data.size >>> ", len(data))

		for index, d := range data {
			var teacher bson.M
			mongo.GetCollection(mongo.TEACHER_COLL).Find(bson.M{"tid": d["teacherId"]}).One(&teacher)
			delete(d, "teacherId")
			delete(d, "_id")
			delete(teacher, "_id")
			d["teacher"] = teacher
			bd, _ := json.Marshal(d)
			dataString = dataString + string(bd) + ","
			if (index + 1) == len(data) {
				tt := d["timestamp"].(float64)
				ts = int(tt)
			}
		}
		dataString = dataString[:len(dataString)-1]
		retString = "{\"top_courses\":[" + topString + "],\"items\":[" + itemString + "],\"data\":[" + dataString + "],\"timestamp\":" + strconv.Itoa(ts) + "}"
	} else {
		log.Println("entering else...")
		//{data:[],timestamp:14xxxxxxx}
		var data []bson.M
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "timestamp": bson.M{"$lte": timestamp}}).Limit(mongo.PAGE_SIZE).All(&data)
		log.Println("data.size >>> ", len(data))
		for index, d := range data {
			var teacher bson.M
			mongo.GetCollection(mongo.TEACHER_COLL).Find(bson.M{"tid": d["teacherId"]}).One(&teacher)
			delete(d, "teacherId")
			delete(d, "_id")
			delete(teacher, "_id")
			d["teacher"] = teacher
			bd, _ := json.Marshal(d)
			dataString = dataString + string(bd) + ","
			if (index + 1) == len(data) {
				tt := d["timestamp"].(float64)
				ts = int(tt)
			}
		}
		if len(dataString) > 0 {
			dataString = dataString[:len(dataString)-1]
		}

		retString = "{\"data\":[" + dataString + "],\"timestamp\":" + strconv.Itoa(ts) + "}"
	}
	log.Println("retString >>> ", retString)
	return retString
}

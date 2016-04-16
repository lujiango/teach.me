package service

import (
	"encoding/json"
	"strconv"

	"gopkg.in/mgo.v2/bson"
	"teach.me/teaching/config"
	"teach.me/teaching/mongo"
	"teach.me/teaching/tlog"
)

func GetCoursesByLocation(location string, timestamp int64) string {
	var retString string = ""
	var topString string = ""
	var itemString string = ""
	var dataString string = ""
	var ts int = 0

	tlog.Info("timestamp : ", timestamp)
	if timestamp == 0 {
		//{top_courses:[],items:[],data:[],timestamp:14xxxxxxx}
		//step 1: get top_courses
		var tops []bson.M
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "isTop": 1}).Limit(config.Gconfig.TopLimit).All(&tops)
		tlog.Info("tops.size : ", len(tops))
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
		tlog.Info("items.size : ", len(items))
		for _, i := range items {
			//remove _id
			delete(i, "_id")
			bi, _ := json.Marshal(i)
			itemString = itemString + string(bi) + ","
		}
		itemString = itemString[:len(itemString)-1]
		// step 3: get data
		var data []bson.M
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "isTop": 0}).Limit(config.Gconfig.PageSize).All(&data)
		tlog.Info("data.size : ", len(data))

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
		//{data:[],timestamp:14xxxxxxx}
		var data []bson.M
		mongo.GetCollection(mongo.COURSE_COLL).Find(bson.M{"location": location, "timestamp": bson.M{"$lte": timestamp}}).Limit(config.Gconfig.PageSize).All(&data)
		tlog.Info("data.size : ", len(data))
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
	return retString
}

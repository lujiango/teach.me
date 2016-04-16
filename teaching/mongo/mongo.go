package mongo

import (
	"strconv"

	"gopkg.in/mgo.v2"
	"teach.me/teaching/config"
)

const (
	TEACHING_DB  = "teaching"
	COURSE_COLL  = "course"
	TEACHER_COLL = "teacher"
	ITEM_COLL    = "item"
)

type Course struct {
	Cid         int     `json:"cid"`
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Thumbnail   string  `json:"thumbnail"`
	TeacherId   string  `json:"teacherId"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Timestamp   int     `json:"timestamp"`
	Total       int     `json:"total"`
	Sign        int     `json:"sign"`
	ItemId      int     `json:"itemId"`
	IsTop       int     `json:"isTop"`
	Price       float64 `json:"price"`
	Address     string  `json:"address"`
	StartTime   int     `json:"startTime"`
	EndTime     int     `json:"endTime"`
	Days        int     `json:"days"`
}

type Teacher struct {
	Tid        int    `bson:"Tid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Age        int    `json:"age"`
	TeachAge   int    `json:"teachAge"`
	Experience string `json:"experience"`
	Sex        int    `json:"sex"`
	IsVerify   int    `json:"isVerify"`
}

type Item struct {
	Iid  int    `json:"iid"`
	Name string `json:"name"`
}

//// query one course by filter condition
//func QueryOne(coll string, filter bson.M) {
//	log.Println("coll: ", coll, "filter: ", filter)
//	course := &Course{}
//	getCollection(coll).Find(filter).One(course)
//}

//// query all courses from coll(collection)
//func QueryAll(coll string, filter string) []Course {
//	log.Println("coll: ", coll, ",filter: ", filter)
//	var courses []Course
//	getCollection(coll).Find(nil).All(&courses)
//	return courses
//}
func GetCollection(coll string) *mgo.Collection {
	session, err := mgo.Dial(config.Gconfig.MgoServer + ":" + strconv.Itoa(config.Gconfig.MgoPort))
	if err != nil {
		panic(err)
	}
	//	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return session.DB(TEACHING_DB).C(coll)
}

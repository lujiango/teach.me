package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	MONGO_DB     = "teaching"
	COURSE_COLL  = "course"
	TEACHER_COLL = "teacher"
)

type Course struct {
	Id_         bson.ObjectId `json:"_id"`
	Name        string        `json:"name"`
	Title       string        `json:"title"`
	Image       string        `json:"image"`
	Thumbnail   string        `json:"thumbnail"`
	TeacherId   string        `json:"teacherid"`
	Description string        `json:"description"`
	Location    string        `json:"location"`
	Timestamp   string        `json:"timestamp"`
	Comment     string        `json:"comment"`
	TotalNumber int           `json:"totalnumber"`
	SignNumber  int           `json:"signnumber"`
}

type Teacher struct {
	Name string `bson:"name"`
}
type Test struct {
	Name string `bson:name`
}

// query one course by filter condition
func QueryOne(coll string, filter string) {
	log.Println("coll: ", coll, "filter: ", filter)
	course := &Course{}
	getCollection(coll).Find(filter).One(course)
}

// query all courses from coll(collection)
func QueryAll(coll string, filter string) []Course {
	log.Println("coll: ", coll, ",filter: ", filter)
	var courses []Course
	getCollection(coll).Find(nil).All(&courses)
	return courses
}
func getCollection(coll string) *mgo.Collection {
	session, err := mgo.Dial("127.0.0.1:20822")
	if err != nil {
		panic(err)
	}
	//	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return session.DB(MONGO_DB).C(coll)

}

package mongo

import (
	"log"
	"fmt"
	"gopkg.in/mgo.v2"
)
type Course struct{
	Name string `bson:"name"`
	Title string `bson:"title"`
	Image string `bson:"image"`
	Thumbnail string `bson:"thumbnail"`
	TeacherId string `bson:"teacherid"`
	Desc stringn `bson:"desc"`
	Location string `bson:"location"`
	Timestamp string `bson:"timestamp"`
	Comment string `bson:"comment"`
}
type Teacher struct{
	Name string `bson:"name"`
}
type Test struct{
	Name string `bson:name`
}

// query one course by filter condition
func QueryOne(filter string) {
	log.Println("collection: ", coll, ", filter: ", filter)
	course := &Course{}
	getCollection().Find(filter).One()
	
}
// query all courses
func QueryAll() []Course{
	var courses []Course
	getCollection().Find(nil).All(&course)
	return courses
}
func getCollection(coll string) *Collection{
	session, err := mgo.Dial("127.0.0.1:20822")
	if err != nil{
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return session.DB("teaching").C(coll)
	
}

session, err := mgo.Dial("server1.example.com,server2.example.com")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
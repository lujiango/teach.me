package http

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"teach.me/teaching/service"
)

const (
	GET_COURSES_BY_LOCATION = "/teaching/course/_get"
)

// get courses by location info.
func GetCourses(w http.ResponseWriter, req *http.Request) {
	location := req.FormValue("location")

	if location == "" {
		io.WriteString(w, "{ret:400100,msg:'location is empty'}")
		return
	}
	timestamp := req.FormValue("timestamp")
	if timestamp == "" {
		timestamp = "0"
	}

	log.Println("location >>> " + location)
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.Panic(err)
		ts = 0
	}
	courses := service.GetCoursesByLocation(location, ts)
	io.WriteString(w, courses)
}
func Router() {
	log.Println(">>> Add router...")
	http.HandleFunc(GET_COURSES_BY_LOCATION, GetCourses)
}
func Start() {
	err := http.ListenAndServe(":10029", nil)
	if err != nil {
		log.Fatal(">>> Teaching start failed...")
	}
}

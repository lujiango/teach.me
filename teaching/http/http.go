package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"teach.me/teaching/service"
)

const (
	GET_COURSES_BY_LOCATION = "/teaching/course/_get"
)

// get courses by location info.
func GetCourses(res http.ResponseWriter, req *http.Request) {
	location, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(location))
	courses := service.GetCoursesByLocation(string(location))
	io.WriteString(res, courses)
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

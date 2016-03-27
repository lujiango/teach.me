package http

import (
	"io"
	"log"
	"net/http"

	"teach.me/teaching/service"
)

const (
	GET_COURSES_BY_LOCATION = "/teaching/course/_get"
)

// get courses by location info.
func GetCourses(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Body)
	courses := service.GetCoursesByLocation("")
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

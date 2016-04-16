package http

import (
	"io"
	"net/http"
	"strconv"

	"teach.me/teaching/config"
	"teach.me/teaching/service"
	"teach.me/teaching/tlog"
)

// get courses by location info.
func GetCourses(w http.ResponseWriter, req *http.Request) {
	location := req.FormValue("location")

	if location == "" {
		io.WriteString(w, LOCATION_IS_EMPTY)
		return
	}
	timestamp := req.FormValue("timestamp")
	if timestamp == "" {
		timestamp = "0"
	}

	tlog.Info("location : ", location)
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		tlog.Error(err)
		ts = 0
	}
	courses := service.GetCoursesByLocation(location, ts)
	io.WriteString(w, courses)
}
func Router() {
	tlog.Info(">>> Add router...")
	http.HandleFunc(TEACHING_INDEX, GetCourses)
}
func Start() {
	err := http.ListenAndServe(":"+strconv.Itoa(config.Gconfig.Port), nil)
	if err != nil {
		tlog.Fatal(">>> Teaching start failed...")
	}
}

package http

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"teach.me/teaching/config"
	"teach.me/teaching/http/req"
	"teach.me/teaching/http/ret"
	"teach.me/teaching/service"
	"teach.me/teaching/tlog"
)

// get courses by location info.
func index(w http.ResponseWriter, req *http.Request) {
	location := req.FormValue("location")

	if location == "" {
		io.WriteString(w, ret.LOCATION_IS_EMPTY)
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
func Html(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/test.html")
	if err != nil {
		tlog.Error(err)
	}
	err = t.Execute(w, "")
	if err != nil {
		tlog.Error(err)
	}
}
func Router() {
	tlog.Info(">>> Add router...")
	http.HandleFunc(req.TEACHING_INDEX, index)
	http.HandleFunc(req.TEACHING_USER_BASE_REG, service.UserBaseReg)
	http.HandleFunc(req.TEACHING_USER_BASE_LOGIN, service.UserBaseLogin)
	http.HandleFunc(req.TEACHING_USER_BASE_PHONE_CHECK, service.UserBasePhoneCheck)
	http.HandleFunc(req.TEACHING_HTML_TEST, Html)
}
func Start() {
	err := http.ListenAndServe(":"+strconv.Itoa(config.Gconfig.Port), nil)
	if err != nil {
		tlog.Fatal(">>> Teaching start failed...")
	}
}

// teaching is a server
package main

import (
	"teach.me/teaching/conf"
	"teach.me/teaching/http"
	"teach.me/teaching/tlog"
)

func main() {
	tlog.Info(">>> Teaching server started...")
	conf.SetConfig()
	tlog.Info(conf.Gconfig)
	http.Router()
	http.Start()
}

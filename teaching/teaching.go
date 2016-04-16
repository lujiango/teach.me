// teaching is a server
package main

import (
	"teach.me/teaching/config"
	"teach.me/teaching/http"
	//	"teach.me/teaching/qiniu"
	"teach.me/teaching/tlog"
)

func main() {

	tlog.Info(">>> Teaching server started...")
	config.SetConfig()
	tlog.Debug(config.Gconfig)

	http.Router()
	http.Start()

}

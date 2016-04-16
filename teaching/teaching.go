// teaching is a server
package main

import (
	"teach.me/teaching/http"
	"teach.me/teaching/tlog"
)

func main() {
	tlog.Info(">>> Teaching server started...")

	http.Router()
	http.Start()
}

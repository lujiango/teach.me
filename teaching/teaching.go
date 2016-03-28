// teaching is a server
package main

import (
	"log"

	"teach.me/teaching/http"
)

func main() {
	log.Println(">>> Teaching server started...")
	http.Router()
	http.Start()
}

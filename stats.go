package main

import (
	"log"
	"net/http"
	"player-stats/models"
)

const httpPort = ":80"

func main() {
	models.HandleList()
	models.HandleAdd()

	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatal(err)
		return
	}
}

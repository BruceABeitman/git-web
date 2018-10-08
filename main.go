package main

import (
	"log"

)

func main() {
	service := core.NewService("https://api.github.com/")
	server := api.BuildHTTPServer(service)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

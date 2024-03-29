package main

import (
	"cafe-management/cmd/api/server"
	"log"
)

// @title          Cafe-Management API
// @version         1.0
// @description     cafe-management
// @contact.name   Veli Ulugut
// @contact.email  veliulugut1@gmail.com
// @host      localhost:8080
// @BasePath  /v1
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	s := server.New(8080)
	err := s.Init()
	if err != nil {
		log.Fatalln(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatalln(err)

	}
}

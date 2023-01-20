package main

import (
	"go-mongo-starter/app"
	_ "go-mongo-starter/docs"
)

// @title Go Mongo Starter API
// @version 0.1
// @description A backend service template for Go and MongoDB
// @contact.name Richard Amare
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	if err := app.SetupAndRunApp(); err != nil {
		panic(err)
	}
}

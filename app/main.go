package main

import (
	"Project/research/sample-gql/config"
	_authController "Project/research/sample-gql/delivery/controllers/auth"
	_bookController "Project/research/sample-gql/delivery/controllers/book"

	"Project/research/sample-gql/delivery/controllers/graph"
	_personController "Project/research/sample-gql/delivery/controllers/person"
	"Project/research/sample-gql/delivery/router"
	"Project/research/sample-gql/repository/auth"
	"Project/research/sample-gql/repository/book"
	"Project/research/sample-gql/repository/person"
	"Project/research/sample-gql/util"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDriver(config)

	//initiate user model
	authRepo := auth.New()
	personRepo := person.New(db)
	bookRepo := book.New(db)

	//initiate user controller
	authController := _authController.New(authRepo)
	personController := _personController.New(personRepo)
	bookController := _bookController.New(bookRepo)

	//create echo http
	e := echo.New()
	client := graph.NewResolver(bookRepo, personRepo)
	srv := router.NewGraphQLServer(client)
	//register API path and controller
	router.RegisterPath(e, authController, personController, bookController, srv)

	// run server
	address := fmt.Sprintf(":%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}

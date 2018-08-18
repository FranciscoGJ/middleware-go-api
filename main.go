package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/negrocu/middleware/apiGin/router"
	"github.com/negrocu/middleware/apiGin/models"
	"github.com/labstack/gommon/log"
	"github.com/negrocu/middleware/apiGin/handlers"
)

const (
	host     = "localhost"
	user     = "negrocu"
	password = "negrocu"
	dbname   = "test"
)

func main() {

	// Set Database
	conf := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := models.NewDB(conf)
	if err != nil {
		log.Panic(err)
	}

	env := &handlers.Env{db}

	r := gin.Default()

	// Set routes
	router.New(r,env)

	// Listen and serve on localhos:8080
	r.Run()
	fmt.Println("Hola mundo")
}

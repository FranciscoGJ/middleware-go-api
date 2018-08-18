package router

import (
	"github.com/gin-gonic/gin"
	"github.com/negrocu/middleware/apiGin/handlers"
)



func New(c *gin.Engine,env *handlers.Env) {

	// User routes
	c.GET("/users", env.UserIndex)
	c.GET("/users/:id", env.UserShow)
	c.POST("/users", env.UserCreate)

	// Sensor routes
	c.GET("/sensors", env.SensorIndex)
}

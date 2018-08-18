package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (env *Env) SensorIndex(c *gin.Context) {

	sensors, err := env.Db.SensorIndex()
	if err != nil {
		c.JSON(http.StatusInternalServerError, 500)
		return
	}
	c.JSON(http.StatusOK,sensors)
}

func (env *Env) SensorShow(c *gin.Context) {
	sensor, err := env.Db.UserShow(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, 500)
		return
	}
	c.JSON(http.StatusOK,sensor)
}

func (env *Env) SensorCreate(c *gin.Context) {
	sensor, err := env.Db.SensorCreate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,sensor)
}
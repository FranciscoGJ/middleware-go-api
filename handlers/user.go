package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (env *Env) UserIndex(c *gin.Context) {

	users, err := env.Db.UserIndex()
	if err != nil {
		c.JSON(http.StatusInternalServerError, 500)
		return
	}
	c.JSON(http.StatusOK,users)
}

func (env *Env) UserShow(c *gin.Context) {
	user, err := env.Db.UserShow(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, 500)
		return
	}
	c.JSON(http.StatusOK,user)
}

func (env *Env) UserCreate(c *gin.Context) {
	user, err := env.Db.UserCreate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,user)
}
package models

import (
		"database/sql"
		_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

type Datastore interface {
	//Users
	UserIndex() ([]*User, error)
	UserShow(c *gin.Context) (*User, error)
	UserCreate(c *gin.Context) (*User, error)

	//Sensors
	SensorIndex() ([]*Sensor, error)
	SensorShow(c *gin.Context) (*Sensor, error)
	SensorCreate(c *gin.Context) (*Sensor, error)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSource string) (*DB,error) {
	db, err := sql.Open("postgres",dataSource)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

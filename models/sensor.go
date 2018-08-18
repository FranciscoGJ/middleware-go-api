package models

import "github.com/gin-gonic/gin"

type Sensor struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Brand string `json:"brand"`
	SKU string `json:"sku"`
	Unit string `json:"unit"`
	UnitDescription string `json:"udescription"`
}

func (db *DB) SensorIndex() ([]*Sensor, error) {
	rows, err := db.Query("SELECT * FROM sensors")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sensors := make([]*Sensor,0)
	for rows.Next() {
		sensor := new(Sensor)
		err := rows.Scan(&sensor.ID,&sensor.Name,&sensor.Description,&sensor.Brand,&sensor.SKU,&sensor.Unit,&sensor.UnitDescription)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors,sensor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sensors, nil
}

func (db *DB) SensorShow(c *gin.Context) (*Sensor, error) {

	id := c.Param("id")
	sqlStatement := `SELECT * FROM sensors where id=$1`
	row := db.QueryRow(sqlStatement, id)
	sensor := new(Sensor)
	err := row.Scan(&sensor.ID,&sensor.Name,&sensor.Description,&sensor.Brand,&sensor.SKU,&sensor.Unit,&sensor.UnitDescription)
	if err != nil {
		return nil, err
	}
	return sensor, nil
}

func (db *DB) SensorCreate(c *gin.Context) (*Sensor, error) {
	sensor := Sensor{}

	if err := c.ShouldBindJSON(&sensor); err != nil {
		return nil, err
	}

	sqlStatement := `INSERT INTO users (id,name,description,brand,sku,unit,unit_description) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, &sensor.ID,&sensor.Name,&sensor.Description,&sensor.Brand,&sensor.SKU,&sensor.Unit,&sensor.UnitDescription).Scan(&sensor.ID)

	if err != nil {
		return nil, err
	}

	return &sensor, nil

}
package models

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int `json:"id"`
	FirstName string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Age int `json:"age"`
}

func (db *DB) UserIndex() ([]*User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User,0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.ID,&user.Age,&user.FirstName,&user.Lastname,&user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users,user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DB) UserShow(c *gin.Context) (*User, error) {

	id := c.Param("id")
	sqlStatement := `SELECT * FROM users where id=$1`
	row := db.QueryRow(sqlStatement, id)
	user := new(User)
	err := row.Scan(&user.ID,&user.Age,&user.FirstName,&user.Lastname,&user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (db *DB) UserCreate(c *gin.Context) (*User, error) {
	user := User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, err
	}

	sqlStatement := `INSERT INTO users (first_name, last_name, email, age) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, &user.FirstName,&user.Lastname,&user.Email,&user.Age).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

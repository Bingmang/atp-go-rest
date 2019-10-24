package models

import (
	"github.com/lib/pq"
	"time"
)

type user struct {
	base
	Account string 			`json:"account"`
	Password string 		`json:"password"`
	LastLogin time.Time 	`json:"last_login"`
}

type Coach struct {
	user
	Majors []string 		`json:"majors"`
	Classes []string 		`json:"classes"`
}

type Student struct {
	user
	Major string			`json:"major"`
	Class string			`json:"class"`
	CoachId int64			`json:"coach_id"`
}

func (c *Coach) Insert() (insertID int64, err error) {
	stmt, err := Conn.Prepare("INSERT INTO coach(account, password, majors, classes, created_time, yn) " +
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return -1, err
	}
	c.Yn = true
	c.CreatedTime = time.Now().UTC()
	res, err := stmt.Exec(c.Account, c.Password, pq.Array(c.Majors), pq.Array(c.Classes), c.CreatedTime, c.Yn)
	if err != nil {
		return -1, err
	}
	insertID, _ = res.LastInsertId()
	return insertID, nil
}

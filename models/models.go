package models

import (
	"github.com/rchenhyy/demo-ginex/db"
	"log"
)

type UserModel struct {
	Id            int64
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
}

func (user *UserModel) Save() int64 {
	result, err := db.Db.Exec("insert into users(email, password) values (?, ?)", user.Email, user.Password)
	if err != nil {
		log.Panicln("user insert error", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

func (user *UserModel) QueryOne() *UserModel {
	u := UserModel{}
	row := db.Db.QueryRow("select * from users where email = ?", user.Email)
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		log.Panicln(err)
	}
	return &u
}

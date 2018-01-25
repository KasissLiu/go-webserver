package models

import (
	"errors"

	"github.com/KasissLiu/go-webserver/dbserver"

	_ "github.com/go-sql-driver/mysql"
)

const (
	tableName = "user"
	connName  = "mysqllocal"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
}

type userModel struct{}

var UserModel userModel

func (u *userModel) GetUserById(id int) (User, error) {
	newUser := User{}
	db, err := dbserver.GetMysql(connName)

	if err != nil {
		return newUser, err
	}
	err = db.QueryRow("select * from "+tableName+" where id = ?", id).Scan(&newUser.Id, &newUser.Name, &newUser.Age, &newUser.Birthday)
	if err != nil {
		return newUser, errors.New("user not found")
	}
	return newUser, nil
}

func (u *userModel) GetUserAll() (users []User) {
	users = make([]User, 0, 10)
	db, err := dbserver.GetMysql(connName)
	rows, err := db.Query("select * from " + tableName)
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Age, &user.Birthday)
		users = append(users, user)
	}

	return
}

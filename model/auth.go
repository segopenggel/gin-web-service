package model

import db "gin-web-app/database"

//import "time"

type Login struct {
	Id        int    `json:"id" form:"id"`
	Access string `json:"access" form:"access"`
	Secret  string `json:"secret" form:"secret"`
}

// fungsi otentikasi
func (a *Login) Auth() (authorized Login, err error) {
	// http://go-database-sql.org/retrieving.html
	row := db.My.QueryRow("SELECT id FROM users WHERE access = ? AND secret = SHA2(?,256)", a.Access, a.Secret)
	err = row.Scan(&authorized.Id)
	if err != nil {
		return
	}
	return
}
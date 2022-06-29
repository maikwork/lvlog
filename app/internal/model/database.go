package model

import "fmt"

type Database struct {
	Type     string
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (db Database) GetDSN() string {
	//postgres://username:password@localhost:5432/database_name
	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable", db.Type, db.Username, db.Password,
		db.Host, db.Port, db.DBName)
	return dsn
}

package database

import (
	"database/sql"
	"example/service/packages/conf"
	"example/service/packages/dataTypes"

	"github.com/go-sql-driver/mysql"
)

type DbServiceStruct struct {
	Database *sql.DB
}

func Initialize() (DbServiceStruct, error) {
	db, err := Connect()
	if err != nil {
		return DbServiceStruct{}, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return DbServiceStruct{}, pingErr
	}
	return DbServiceStruct{Database: db}, nil
}

func dsn() string {
	c := conf.GetDbCredentials()
	cfg := mysql.Config{
		User:   c.Username,
		Passwd: c.Password,
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: c.DbName,
	}
	return cfg.FormatDSN()
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
	return db, err
}

func (dss *DbServiceStruct) GetById(id string) (dataTypes.UserData, error) {
	var user dataTypes.UserData
	tx, err := dss.Database.Begin()
	defer tx.Rollback()
	if err != nil {
		return user, err
	}
	rows, err := tx.Query("SELECT id, info FROM users WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		return user, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Info); err != nil {
			return user, err
		}
	}
	return user, nil
}

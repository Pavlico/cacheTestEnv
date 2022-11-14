package service

import (
	"example/service/packages/dataTypes"
	"example/service/packages/database"
	"log"
)

func GetUserData(id string) *dataTypes.UserData {
	db, err := database.Initialize()
	if err != nil {
		log.Println(err)
		return nil
	}
	userData, err := db.GetById(id)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &userData
}

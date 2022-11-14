package handler

import (
	"example/service/packages/conf"
	"example/service/packages/dataTypes"
	"example/service/packages/http/service"
	"net/http"
)

func GetUserById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var emptyUser dataTypes.UserData
		var status int = http.StatusOK
		id := r.FormValue(conf.UserIdKey)
		userData := service.GetUserData(id)
		result := userData.Info
		if *userData == emptyUser {
			status = http.StatusNotFound
			result = conf.NotFoundMessage
		}
		w.WriteHeader(status)
		w.Write([]byte(result))
	}
}

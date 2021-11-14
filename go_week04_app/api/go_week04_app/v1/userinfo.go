package v1

import (
	"encoding/json"
	"net/http"
)

type UserInfoDTO struct {
	Name     string
	UserName string
}

type UserHandler interface {
	ShowUserInfo(id string) *UserInfoDTO
}

func BuildUserInfoHandler(use UserHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		id := request.Form.Get("id")
		userinfo := use.ShowUserInfo(id)
		bytes, _ := json.Marshal(userinfo)
		_, _ = writer.Write(bytes)
	}
}

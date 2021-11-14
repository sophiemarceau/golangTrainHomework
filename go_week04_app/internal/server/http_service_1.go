package server

import (
	v1 "go_week04_app/api/go_week04_app/v1"
	"go_week04_app/internal/conf"
	"go_week04_app/internal/service"
	"net/http"
	"strconv"
)

type HttpServer1 http.Server

func InitHttpServer1(config *conf.Config, user *service.UserServcie) *HttpServer1  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", v1.BuildUserInfoHandler(user))
	return &HttpServer1{
		Addr: ":" + strconv.Itoa(config.Server1Port),
		Handler: mux,
	}
}

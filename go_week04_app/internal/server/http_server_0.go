package server

import (
	v1 "go_week04_app/api/go_week04_app/v1"
	"go_week04_app/internal/conf"
	"go_week04_app/internal/service"
	"net/http"
	"strconv"
)

type HttpServer0 http.Server

func InitHttpServer0(config *conf.Config, user *service.UserServcie) *HttpServer0 {
	mux := http.NewServeMux()
	mux.HandleFunc("/", v1.BuildUserInfoHandler(user))
	return &HttpServer0{
		Addr: ":" + strconv.Itoa(config.Server0Port),
		Handler: mux,
	}
}

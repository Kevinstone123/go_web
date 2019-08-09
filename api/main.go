package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"go_web/api/handlers"
)

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()

	// 用户注册
	router.POST("/user", handlers.CreateUser)

	// 用户登陆
	router.POST("/user/:user_name", handlers.Login)

	return router
}

func main() {
	r := RegisterHandles()
	http.ListenAndServe(":8888", r)
}
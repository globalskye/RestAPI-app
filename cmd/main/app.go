package main

import (
	"fmt"
	"github.com/globalskye/RestAPI-app.git/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("register user handler")

	handler := user.NewHandler()
	handler.Register(router)
	start(router)

}
func start(router *httprouter.Router) {
	log.Println("start application")

	v := "1234"

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", v))
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Println("server is listening on port", v)
	log.Fatalln(server.Serve(listener))
}

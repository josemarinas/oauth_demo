package main

import (
	"net/http"
	"oauth_demo/handlers"
	"oauth_demo/config"
)
func main() {
	config.Init()
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/error", handlers.Error)
	http.HandleFunc("/callback", handlers.Callback)
	http.ListenAndServe(":3000", nil)
}


package main

import (
	"CSA_1101_-_Authentication/back_end"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	err := back_end.Connect()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/validate", back_end.Register)
	mux.HandleFunc("/authenticate", back_end.Login)
	mux.HandleFunc("/reset", back_end.Forgot)
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server started at port 8080")
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

package main

import (
	"CSA_1101_-_Authentication/back_end"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/validate", back_end.Request)
	fmt.Println("Server started at port 8080")
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

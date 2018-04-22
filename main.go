package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func main() {

	router := httprouter.New()

	fmt.Println("initializing routes..")
	router.GET("/search", searchHandler)

	fmt.Println("listening and serving..")

	l, err := net.Listen("tcp4", ":8181")

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(l, router))
}

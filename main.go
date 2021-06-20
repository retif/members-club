package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
)

var club = Club{}
var templates = template.Must(template.New("homepage.html").Funcs(template.FuncMap{"add": func(x,y int) int {return x+y}}).ParseFiles("templates/homepage.html"))

func main() {
	initRoutes()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("WebServer is running: http://localhost:8080")

	log.Fatal(http.Serve(l, nil))
}

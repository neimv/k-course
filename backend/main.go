package main

import (
	"flag"
	"fmt"
)

var (
	host = flag.String("h", "localhost", "Host to start web server")
	port = flag.Int("p", 3000, "port to access web server")
)

func main() {
	url := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Println("Starting server")
	db := GetDB()
	db.AutoMigrate(&ToDoORM{})

	router := CreateRouter()

	router.Logger.Fatal(router.Start(url))
}

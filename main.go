package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/starlingvibes/Simple-CRUD-MonGo/configs"
	"github.com/starlingvibes/Simple-CRUD-MonGo/routes"
)

func main() {
    router := mux.NewRouter()

	configs.ConnectDB()
	routes.UserRoute(router)

	fmt.Println("Server is running on port 1337")
    log.Fatal(http.ListenAndServe(":1337", router))
}
package main

import (
	"ApplicationGo/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Conexão feita!")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

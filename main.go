package main

import (
	"fmt"
	"net/http"
	"text/template"

)



var temp = template.Must(template.ParseGlob("templates/*.html")) // Carrega todos os templates presentes no programa

func main() {
	fmt.Println("Conexão feita!")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func inicio(w http.ResponseWriter, r *http.Request) {
	pagina, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "No se pudo cargar la pagina", http.StatusInternalServerError)
		fmt.Println("Error:", err)
		return
	}

	pagina.Execute(w, nil)

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", inicio)
	http.HandleFunc("/productos", productos)
	fmt.Println("Servidor inciciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
func productos(w http.ResponseWriter, r *http.Request) {
	pagina, _ := template.ParseFiles("templates/productos.html")
	pagina.Execute(w, nil)
}

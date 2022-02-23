package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func res(w http.ResponseWriter, r *http.Request) {

	template.ParseFiles("templates/index.html")
	fmt.Fprintf(w, " Где шаблон ???!")
}

func main() {

	fmt.Println("Запущен!")
	http.HandleFunc("/firstApi", res)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Не стартануло!")
	}

}

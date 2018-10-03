package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("web/index.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := TodoPageData{
			PageTitle:"My Todo List",
			Todos:[]Todo{
				{Title:"Task 1", Done:false},
				{Title:"Task 2", Done:true},
				{Title:"Task 3", Done:false},
			},
		}
		tmpl.Execute(writer, data)
	})
	log.Println("Starting web server at 127.0.0.1:8080...")
	http.ListenAndServe(":8080", nil)
}

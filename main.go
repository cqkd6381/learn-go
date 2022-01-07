package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/index1.html", indexHandler)
	http.HandleFunc("/index2.html", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("page not exist")
	})
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	http.Handle("/plugins/", http.StripPrefix("/plugins/", http.FileServer(http.Dir("plugins"))))
	http.Handle("/bower_components/", http.StripPrefix("/bower_components/", http.FileServer(http.Dir("bower_components"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("pages"))))
	http.ListenAndServe(":9090", nil)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("index1.html"))
	t.Execute(writer, "")
}

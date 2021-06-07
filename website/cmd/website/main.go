package main

import (
//	"fmt"
	"net/http"
//	"html/template"
)


func main(){
	http.HandleFunc("/", home)
	http.HandleFunc("/notebook/", notebook)
	http.HandleFunc("/taskbook/", taskbook)
//	http.HandleFunc("calendar/", calendar)
	fileServer := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.ListenAndServe(":8888", nil)
}

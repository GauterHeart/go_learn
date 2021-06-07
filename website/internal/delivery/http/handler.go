package main

import (
//	"fmt"
	"net/http"
	"html/template"
)


func home(w http.ResponseWriter, r *http.Request){
	tmp, _ := template.ParseFiles("./web/html/base.html", "./web/html/home.html")
	tmp.ExecuteTemplate(w, "base", "")
}

func notebook(w http.ResponseWriter, r *http.Request){
	tmp, _ := template.ParseFiles("./web/html/base.html", "./web/html/notebook.html")
	tmp.ExecuteTemplate(w, "base", "")
}

func taskbook(w http.ResponseWriter, r *http.Request){
	tmp, _ := template.ParseFiles("./web/html/base.html", "./web/html/taskbook.html")
	tmp.ExecuteTemplate(w, "base", "")
}


package main

import (
//	"fmt"
	"html/template"
//	"log"
	"net/http"
)

type User struct {
	name string
	age uint16
}

func home(w http.ResponseWriter, r *http.Request){
	sam := User{"Sam", 19}
	tml, _ := template.ParseFiles("html/home.html", "html/base.html")
	tml.ExecuteTemplate(w, "home", sam)
}


func main(){
	http.HandleFunc("/", home)
	http.ListenAndServe(":8888", nil)

}

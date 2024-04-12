package main

import(
	"net/http"
	"html/template"
) 
 func (d DB) http(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	   tmpl, _ := template.ParseFiles("./tmpl/index.html")
	   tmpl.Execute(w, tmpl)

    })

 }

package main

import(
	"net/http"
	"html/template"
) 
 func (d DB) http(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	   tmpl, _ := template.ParseFiles("./tmpl/sf.html")
	   tmpl.Execute(w, tmpl)

    })
    http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request){
		tmpl, _ := template.ParseFiles("./tmpl/sd.html")
		tmpl.Execute(w, tmpl)
 
	 })
	 http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request){
		tmpl, _ := template.ParseFiles("./tmpl/ls.html")
		tmpl.Execute(w, tmpl)
 
	 })
	 http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request){
		tmpl, _ := template.ParseFiles("./tmpl/os.html")
		tmpl.Execute(w, tmpl)
 
	 })
	 http.HandleFunc("/d", func(w http.ResponseWriter, r *http.Request){
		tmpl, _ := template.ParseFiles("./tmpl/main.html")
		tmpl.Execute(w, tmpl)
 
	 })
 }

package main

import (
	"html/template"
	"net/http"
)

func (d DB) http() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tmpl/index.html")
		tmpl.Execute(w, tmpl)

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tmpl/frO.html")
		tmpl.Execute(w, tmpl)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tmpl/frT.html")
		tmpl.Execute(w, tmpl)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tmpl/frTh.html")
		tmpl.Execute(w, tmpl)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tmpl/frFo.html")
		tmpl.Execute(w, tmpl)
	})
}

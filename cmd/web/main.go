package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func init() {
	t := template.New("").Funcs(template.FuncMap{"unescaped": func(x string) interface{} { return template.HTML(x) }})
	var err error
	templates, err = t.ParseGlob("templates/*.html")
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", item)
	r.HandleFunc("/{id}", item)
	r.HandleFunc("/login/", login)
	r.HandleFunc("/search/", search)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", r))
}

func item(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		err := templates.ExecuteTemplate(w, "index.html", homeData)
		if err != nil {
			http.Error(w, "could not execute template", http.StatusInternalServerError)
		}
		return
	}
	if item, ok := itemsData[id]; ok {
		err := templates.ExecuteTemplate(w, "index.html", item)
		if err != nil {
			http.Error(w, "could not execute template", http.StatusInternalServerError)
		}
		return
	}
	err := templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		http.Error(w, "could not execute template", http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "could not execute template", http.StatusInternalServerError)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "could not execute template", http.StatusInternalServerError)
	}
}

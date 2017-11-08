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
		templates.ExecuteTemplate(w, "index.html", homeData)
		return
	}
	if item, ok := itemsData[id]; ok {
		templates.ExecuteTemplate(w, "index.html", item)
		return
	}
	templates.ExecuteTemplate(w, "404.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

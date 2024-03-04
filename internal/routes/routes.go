package routes

import (
	"html/template"
	"log"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// resolve requests to staticfiles (css, imgs, ...)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/minion", minionHandler)

	return mux
}

func minionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/fragments/minion.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Panicln("minionHandler: error executing template")
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Panicln("rootHandler: error executing template")
	}
}

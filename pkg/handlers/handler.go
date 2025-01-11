package handlers

import (
	"html/template"
	//"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func Main() {

	mux := http.NewServeMux()
	//rh := http.RedirectHandler("http://example.org", 307)
	//mux.Handle("/foo", rh)
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/home", handleMainPage)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/register", handleRegister)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)

	}
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Print(err)
	}
	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Print(err)
	}
}
func handleLogin(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "login.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Print(err)
	}
	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Print(err)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "register.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Print(err)
	}
	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Print(err)
	}
}

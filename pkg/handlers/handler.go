package handlers

import (
	"database/sql"
	"strconv"

	//"strings"

	"encoding/json"
	dbPkg "github.com/AlexeyChudov/todoApp/pkg/Database"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Task struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Completed    bool      `json:"completed"`
	CreationDate time.Time `json:"creation_date"`
	DueDate      time.Time `json:"due_date"`
	Priority     string    `json:"priority"`
}
type TaskWithStrDates struct {
	Task
	StrCreationDate string `json:"str_creation_date"`
	StrDueDate      string `json:"str_due_date"`
}

var db *sql.DB

func Main() {
	connStr := "user=smash password=smash host=localhost port=5432 dbname=todoAppDB sslmode=disable"
	db = dbPkg.Init("postgres", connStr)

	router := mux.NewRouter()
	//rh := http.RedirectHandler("http://example.org", 307)
	//router.Handle("/foo", rh)
	//var dir string

	//flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	//flag.Parse()
	fs := http.FileServer(http.Dir("./static/"))
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/home", renderMainPage).Methods("GET")
	router.HandleFunc("/home/tasks", handleTasks).Methods("GET")
	router.HandleFunc("/home/tasks/", handleNewTask).Methods("POST")
	router.HandleFunc("/home/tasks/{id}", handleDeleteTask).Methods("DELETE")
	router.HandleFunc("/home/tasks/{id}", handleUpdateTask).Methods("PUT")
	router.HandleFunc("/home/tasks/status/{id}", handleUpdateStatus).Methods("PUT")
	router.HandleFunc("/login", handleLogin).Methods("GET")
	router.HandleFunc("/register", handleRegister).Methods("GET")

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)

	}
}
func renderMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := renderTemplate("layout.html", "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
	//tasks, err := SelectAllTasks("")
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}

}
func renderTemplate(layout, templateName string) (*template.Template, error) {
	lp := filepath.Join("templates", layout)
	fp := filepath.Join("templates", templateName)
	log.Println("Layout template", lp, "extended template: ", fp)
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Print(err)
	}
	log.Println(tmpl.Tree)
	return tmpl, err
}

func handleTasks(w http.ResponseWriter, r *http.Request) {

	sortMethod := r.URL.Query().Get("sort")
	filterMethod := r.URL.Query().Get("filter")
	var page int = 1
	if r.URL.Query().Get("page") != "" {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			log.Print(err)
		}
	}

	tasks, err := SelectAllTasks(sortMethod, filterMethod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
	data, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		log.Fatal(err)
	}

}

func handleNewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	dueDate, err := time.Parse("2006-01-02", r.FormValue("due_date"))
	if err != nil {
		log.Fatal(err)
	}
	task := Task{
		Title:        r.FormValue("title"),
		Description:  r.FormValue("description"),
		CreationDate: time.Now(),
		DueDate:      dueDate,
		Priority:     r.FormValue("priority"),
	}
	strCreationDate, err := task.CreationDate.MarshalText()

	strDueDate := task.DueDate.Format("2006-01-02")
	err = dbPkg.InsertStatement(db, task.Title, task.Description, string(strCreationDate), strDueDate, task.Priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	taskId, err := extractIdFromURL(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
	}

	dueDate, err := time.Parse("2006-01-02", r.FormValue("dueDate"))
	if err != nil {
		log.Fatal(err)
	}
	task := Task{
		Id:           taskId,
		Title:        r.FormValue("title"),
		Description:  r.FormValue("description"),
		CreationDate: time.Now(),
		DueDate:      dueDate,
		Priority:     r.FormValue("priority"),
	}

	err = dbPkg.UpdateStatement(db, taskId, task.Title, task.Description,
		task.DueDate.Format("2006-01-02"), task.Priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func handleUpdateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	status := r.URL.Query().Get("completed")
	taskId, err := extractIdFromURL(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
	}
	err = dbPkg.UpdateStatus(db, taskId, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
func handleDeleteTask(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	if request.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	taskId, err := extractIdFromURL(request, "id")

	err = dbPkg.DeleteTaskById(db, taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	//w.WriteHeader(http.StatusNoContent)
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

package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	dbPkg "github.com/AlexeyChudov/todoApp/pkg/Database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"
)

func SelectAllTasks(sortMethod, filter string) ([]TaskWithStrDates, error) {
	allowedSortMethods := []string{"due_date", "creation_date"}
	allowedFilterMethods := []string{"completed", "pending"}
	var rows *sql.Rows
	var err error
	var query string
	switch {
	case slices.Contains(allowedSortMethods, sortMethod) && slices.Contains(allowedFilterMethods, filter):
		if filter == "pending" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = false ORDER BY %s ASC", sortMethod)
		} else if filter == "completed" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = true ORDER BY %s ASC", sortMethod)
		}

	case !slices.Contains(allowedSortMethods, sortMethod) && slices.Contains(allowedFilterMethods, filter):
		if filter == "pending" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = false ORDER BY id")
		} else if filter == "completed" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = true ORDER BY id")
		}
	case slices.Contains(allowedSortMethods, sortMethod) && !slices.Contains(allowedFilterMethods, filter):
		query = fmt.Sprintf("SELECT * FROM tasks ORDER BY %s ASC", sortMethod)
	default:
		query = fmt.Sprintf("SELECT * FROM tasks ORDER BY id", sortMethod)

	}
	rows, err = dbPkg.SelectQuery(db, query)
	if err != nil {
		log.Fatal(err)

	}

	tasks := make([]TaskWithStrDates, 0)
	for rows.Next() {
		task := new(TaskWithStrDates)
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreationDate, &task.DueDate, &task.Priority, &task.Completed)
		if err != nil {
			log.Fatal(err)
		}
		task.StrCreationDate = task.CreationDate.Format("2006-01-02")
		task.StrDueDate = task.DueDate.Format("2006-01-02")
		tasks = append(tasks, *task)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks, err
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Обработка preflight-запроса
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func extractIdFromURL(r *http.Request, item string) (int, error) {
	id := mux.Vars(r)[item]
	if id == "" {
		err := errors.New("id is empty")
		return 0, err
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	return idInt, err
}
func taskDateToStringDate(t time.Time) string {
	return t.Format("2006-01-02")
}

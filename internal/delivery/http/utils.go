package Server

import (
	"database/sql"
	"errors"
	"fmt"
	dbPkg "github.com/AlexeyChudov/todoApp/pkg/Repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"
)

type TasksData struct {
	tasks      []TaskStrDates
	pagesCount int
}

func SelectTasks(sortMethod, filter string, tasksCount, offset int) ([]TaskStrDates, error) {
	allowedSortMethods := []string{"due_date", "creation_date"}
	allowedFilterMethods := []string{"all", "completed", "pending"}
	var rows *sql.Rows
	var err error
	var query string
	switch {
	case slices.Contains(allowedSortMethods, sortMethod) && slices.Contains(allowedFilterMethods, filter):
		if filter == "pending" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = false ORDER BY %s ASC", sortMethod)
		} else if filter == "completed" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = true ORDER BY %s ASC", sortMethod)
		} else {
			query = fmt.Sprintf("SELECT * FROM tasks ORDER BY %s ASC", sortMethod)
		}

	case !slices.Contains(allowedSortMethods, sortMethod) && slices.Contains(allowedFilterMethods, filter):
		if filter == "pending" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = false ORDER BY id")
		} else if filter == "completed" {
			query = fmt.Sprintf("SELECT * FROM tasks WHERE status = true ORDER BY id")
		} else {
			query = fmt.Sprintf("SELECT * FROM tasks ORDER BY id ASC")
		}
	case slices.Contains(allowedSortMethods, sortMethod) && !slices.Contains(allowedFilterMethods, filter):
		query = fmt.Sprintf("SELECT * FROM tasks ORDER BY %s ASC", sortMethod)
	default:
		query = fmt.Sprintf("SELECT * FROM tasks ORDER BY id")

	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d;", tasksCount, offset)
	log.Println(query)
	rows, err = dbPkg.SelectQuery(db, query)
	if err != nil {
		log.Fatal(err)

	}

	tasks := make([]TaskStrDates, 0)
	for rows.Next() {
		task := new(TaskStrDates)
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

func pagination(arr []TaskStrDates, itemsPerPage, pageWanted int) ([]TaskStrDates, error) {
	var currentPosition int
	if pageWanted > (len(arr)/itemsPerPage)+1 {
		return nil, errors.New("page wanted > pages count")
	}
	for currentPage := 0; currentPage < (len(arr)/itemsPerPage)+1; currentPage += 1 {
		if currentPage+1 == pageWanted {
			currentPosition = currentPage * itemsPerPage
			break
		}
	}
	if len(arr) < itemsPerPage {
		return arr[currentPosition:len(arr)], nil
	}
	return arr[currentPosition : currentPosition+itemsPerPage], nil
}

func getTasks(db *sql.DB, sortMethod, filterMethod string, tasksCount, offset int) (TasksData, error) {
	tasks, err := SelectTasks(sortMethod, filterMethod, tasksCount, offset)
	if err != nil {
		log.Fatal(err)
	}

	tasksPerPage := 5

	totalTasks, err := dbPkg.GetRowsCount(db)
	if err != nil {
		log.Println(err)
	}
	divCheck := 0
	if (totalTasks % tasksPerPage) > 0 {
		divCheck = 1
	}

	//log.Println(totalTasks)
	totalPages := (totalTasks / tasksPerPage) + divCheck

	//currentPageTasks, err := pagination(tasks, tasksPerPage, page)
	if err != nil {
		log.Println(err)
	}
	data := TasksData{
		tasks:      tasks,
		pagesCount: totalPages,
	}
	return data, err
}

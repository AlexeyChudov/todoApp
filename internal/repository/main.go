package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Init(driver, connString string) *sql.DB {
	db, err := sql.Open(driver, connString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SelectQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	return rows, err
}
func DeleteTaskById(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

func InsertStatement(db *sql.DB, title, description, creation_date, due_date, priority string) error {
	_, err := db.Exec("INSERT INTO tasks (title, description, creation_date, due_date, priority) VALUES ( $1, $2, $3, $4, $5 )",
		title, description, creation_date, due_date, priority)
	return err
}

func UpdateStatement(db *sql.DB, id int, title, description, dueDate, priority string) error {
	_, err := db.Exec("UPDATE  tasks SET title=$1, description=$2, due_date=$3::DATE, priority=$4 WhERE id=$5",
		title, description, dueDate, priority, id)

	return err
}
func UpdateStatus(db *sql.DB, id int, status string) error {
	_, err := db.Exec("UPDATE  tasks SET status=$1 WhERE id=$2",
		status, id)

	return err
}

func GetRowsCount(db *sql.DB) (int, error) {

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	return count, err
}

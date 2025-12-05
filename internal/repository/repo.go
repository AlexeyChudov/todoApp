package repository

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

const (
	errUrlExists string = "URL already exists"
)

type storageErrors struct {
	errUrlExists error
}

func NewStorage(connStr string) (*Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	storage := &Storage{db}
	err = storage.initStorage()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage: %w", err)
	}
	return storage, nil
}

func (storage *Storage) initStorage() error {
	usersStmnt, err := storage.db.Prepare(`CREATE TABLE IF NOT EXISTS public.users (
			id integer NOT NULL,
			first_name character varying(50) NOT NULL,
			last_name character varying(50) NOT NULL,
			email character varying(100),
			country character varying(100) NOT NULL,
			date_of_birth date NOT NULL,
			car_id bigint,
			CONSTRAINT date_later_1990 CHECK ((date_of_birth >= '1990-01-01'::date))
		);`)
	tasksStmnt, err := storage.db.Prepare(`CREATE TABLE IF NOT EXISTS public.users (
			id integer NOT NULL,
			first_name character varying(50) NOT NULL,
			last_name character varying(50) NOT NULL,
			email character varying(100),
			country character varying(100) NOT NULL,
			date_of_birth date NOT NULL,
			car_id bigint,
			CONSTRAINT date_later_1990 CHECK ((date_of_birth >= '1990-01-01'::date))
		);`)
	if err != nil {
		return fmt.Errorf("failed to prepare init statements: %w", err)
	}
	_, err = usersStmnt.Exec()
	_, err = tasksStmnt.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute init statement: %w", err)
	}
	return nil
}

func (storage *Storage) InsertTask(title, description, creation_date, due_date, priority string) (int64, error) {
	const op = "storage.postgres.InsertTask()"
	res, err := storage.db.Exec("INSERT INTO tasks (title, description, creation_date, due_date, priority) VALUES ( $1, $2, $3, $4, $5 )",
		title, description, creation_date, due_date, priority)
	if err != nil {
		if postgresErr, ok := err.(pq.PGError); ok && postgresErr.Error() == "23505" {
			return 0, fmt.Errorf("%s, %s", op, errUrlExists)
		}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	return id, nil
}

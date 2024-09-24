package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	md "github.com/n-nourdine/todo/models"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Init() {
	db := New()
	defer db.Close()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			status BOOLEAN DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

}

func Add(db *sql.DB, todo md.TodoModel) (int, error) {
	stmt, err := db.Prepare("insert into todos(title, status, created_at) values(?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(todo.Title, false, time.Now())
	if err != nil {
		return -1, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), err
}

func GetAll(db *sql.DB) ([]md.Todo, error) {
	rows, err := db.Query("SELECT id, title, status FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []md.Todo
	for rows.Next() {
		var todo md.Todo
		err := rows.Scan(&todo.TodoId, &todo.Title, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func GetById(db *sql.DB, id int) (*md.Todo, error) {
	var todo md.Todo
	err := db.QueryRow("SELECT id, title, status FROM todos WHERE id = ?", id).Scan(&todo.TodoId, &todo.Title, &todo.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

func Delete(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM todos WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}

func Update(db *sql.DB, todo md.Todo) error {
	stmt, err := db.Prepare("UPDATE todos SET title=?, status=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.Title, todo.Status, todo.TodoId)
	return err
}

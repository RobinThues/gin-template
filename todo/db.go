package todo

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlite3      = "sqlite3"
	createQuery  = "CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, text TEXT, isDone BOOLEAN)"
	updateQuery  = "REPLACE INTO todos (id, text, isDone) VALUES (?, ?, ?)"
	findQuery    = "SELECT id, text, isDone FROM todos WHERE id = ?"
	findAllQuery = "SELECT id, text, isDone FROM todos"
	insertQuery  = "INSERT INTO todos (text, isDone) VALUES (?, ?)"
)

var TodoDb TodoDatabase

func init() {
	TodoDb = NewSqliteTodoDatabase("./todos.db")
}

type TodoDatabase interface {
	SaveTodo(todo Todo)
	FindTodo(id int) *Todo
	FindTodos() Todos
	InsertTodo(todo Todo)
}

type SqliteTodoDatabase struct {
	db *sql.DB
}

func NewSqliteTodoDatabase(file string) *SqliteTodoDatabase {
	database, err := sql.Open(sqlite3, file)
	if err != nil {
		fmt.Println(err)
	}

	createStatement, err := database.Prepare(createQuery)
	if err != nil {
		fmt.Println("onCreate", err)
	}
	_, err = createStatement.Exec()
	if err != nil {
		fmt.Println(err)
	}

	sqliteDb := SqliteTodoDatabase{
		db: database,
	}
	return &sqliteDb
}

func (db *SqliteTodoDatabase) SaveTodo(todo Todo) {
	insertStatement, err := db.db.Prepare(updateQuery)
	if err != nil {
		fmt.Println(err)
	}
	_, err = insertStatement.Exec(todo.Id, todo.Text, todo.IsDone)
	if err != nil {
		fmt.Println(err)
	}
}

func (db *SqliteTodoDatabase) FindTodo(id int) *Todo {
	rows, _ := db.db.Query(findQuery, id)
	var foundId int
	var text string
	var isDone bool
	for rows.Next() {
		rows.Scan(&foundId, &text, &isDone)
	}

	return &Todo{
		Id:     foundId,
		Text:   text,
		IsDone: isDone,
	}
}

func (db *SqliteTodoDatabase) FindTodos() Todos {
	rows, err := db.db.Query(findAllQuery)
	if err != nil {
		fmt.Println(err)
	}
	var todos Todos
	for rows.Next() {
		todo := Todo{
			Id:     0,
			Text:   "",
			IsDone: false,
		}
		err = rows.Scan(&todo.Id, &todo.Text, &todo.IsDone)
		if err != nil {
			fmt.Println(err)
		}
		todos = append(todos, &todo)
	}
	return todos
}

func (db *SqliteTodoDatabase) InsertTodo(todo Todo) {
	insertStatement, err := db.db.Prepare(insertQuery)
	if err != nil {
		fmt.Println(err)
	}
	_, err = insertStatement.Exec(todo.Text, todo.IsDone)
	if err != nil {
		fmt.Println(err)
	}
}

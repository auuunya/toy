package model

import (
	"time"
)

type Todo struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Is_success bool   `json:"is_success"`
	Data       string `json:"date"`
	User_id    int    `json:"user_id"`
}

type Todos []*Todo

func GetAllTodo() ([]*Todo, error) {
	todos := []*Todo{}
	rows, err := Db.Query("select id,title,content,is_success,data,user_id from t_todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		todo := new(Todo)
		_ = rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Is_success, &todo.Data, &todo.User_id)
		todos = append(todos, todo)
	}
	return todos, nil
}
func GetUserTodo(user_id int) ([]*Todo, error) {
	todos := []*Todo{}
	rows, err := Db.Query("select id,title,content,is_success,data from t_todo where user_id=?", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		todo := new(Todo)
		_ = rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Is_success, &todo.Data)
		todos = append(todos, todo)
	}
	return todos, nil
}
func GetTodo(id int) (*Todo, error) {
	todo := new(Todo)
	err := Db.QueryRow("select id,title,content,is_success,data,user_id from t_todo where id=?", id).
		Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Is_success, &todo.Data, &todo.User_id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (user *User) CreateTodo(title, content string) error {
	stmt, err := Db.Prepare("insert into t_todo (title,content,data,user_id) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, content, TimeFormat(time.Now()), user.Id)
	if err != nil {
		return err
	}
	return nil
}
func (todo *Todo) UpdateTodo() error {
	stmt, err := Db.Prepare("update t_todo set is_success=true where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(todo.Id)
	if err != nil {
		return err
	}
	return nil
}

func (todo *Todo) DeleteTodo() error {
	stmt, err := Db.Prepare("delete from t_todo where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(todo.Id)
	if err != nil {
		return err
	}
	return nil
}

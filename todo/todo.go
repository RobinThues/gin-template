package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Todo struct {
	Id		int
	Text	string
	IsDone	bool
}

type Todos = []*Todo

var todos Todos
var counter int

func newTodo(text string) *Todo {
	t := Todo{
		Id:     counter,
		Text:   text,
		IsDone: false,
	}
	counter += 1
	return &t
}

func GetTodos(c *gin.Context) {
	c.Set("data", gin.H{
		"todos": todos,
		"greeting": "Please add a todo below",
	})
	c.Set("tmpl", "todo/list.tmpl")
}

func MarkTodoAsDone(c *gin.Context) {
	parameterId, _ := strconv.Atoi(c.Param("id"))
	for _, t := range todos {
		if t.Id == parameterId {
			t.IsDone = true
		}
	}
	c.Redirect(http.StatusSeeOther, "/todos")
}

func CreateTodo(c *gin.Context) {
	text := c.PostForm("text")
	todos = append(todos, newTodo(text))
	c.Redirect(http.StatusSeeOther, "/todos")
}
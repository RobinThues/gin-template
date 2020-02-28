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

func GetTodos(c *gin.Context) {
	todos := TodoDb.FindTodos()

	c.Set("data", gin.H{
		"todos": todos,
		"greeting": "Please add a todo below",
	})
	c.Set("tmpl", "todo/list.tmpl")
}

func MarkTodoAsDone(c *gin.Context) {
	parameterId, _ := strconv.Atoi(c.Param("id"))

	t := TodoDb.FindTodo(parameterId)
	t.IsDone = true
	TodoDb.SaveTodo(*t)

	c.Redirect(http.StatusSeeOther, "/todos")
}

func CreateTodo(c *gin.Context) {
	text := c.PostForm("text")
	TodoDb.InsertTodo(Todo{
		Id:     0,
		Text:   text,
		IsDone: false,
	})
	c.Redirect(http.StatusSeeOther, "/todos")
}

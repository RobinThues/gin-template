package main

import (
	"fmt"
	"github.com/robinthues/gin-template/todo"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)



func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("./templates/**/*")
	r.Static("/assets", "./assets")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	r.Use(FormatMiddleware())

	// add a testing todo
	todo.TodoDb.SaveTodo(todo.Todo{
		Id:     0,
		Text:   "hello world",
		IsDone: false,
	})

	r.GET("/todos", todo.GetTodos)
	r.GET("/todo/:id/done", todo.MarkTodoAsDone)
	r.POST("/todos", todo.CreateTodo)
	r.GET("/ping", ping)

	if err :=r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func FormatMiddleware() gin.HandlerFunc {
	// initialize middleware
	return func (c *gin.Context) {
		c.Next()
		queriedFormat := c.Query("format")
		data, _ := c.Get("data")

		if queriedFormat == "json" {
			c.JSON(http.StatusOK, data)
		} else {
			tmplName, ok := c.Get("tmpl")
			if ok {
				c.HTML(http.StatusOK, tmplName.(string), data)
			}
		}
	}
}


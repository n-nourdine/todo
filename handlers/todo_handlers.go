package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/n-nourdine/todo/database"
	md "github.com/n-nourdine/todo/models"
)

type Template struct {
	template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func addTodo(c echo.Context) error {
	todo := md.Todo{}
	err := c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong data filful")
	}

	if strings.TrimSpace(todo.Title) == "" {
		return c.JSON(http.StatusBadRequest, "title can not be empty!")
	}

	conn := db.New()
	defer conn.Close()

	todoM := md.TodoModel{
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: time.Now(),
	}

	id, err := db.Add(conn, todoM)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to add todo!")
	}
	// return c.Render(http.StatusOK, "index", []md.Todo{
	// 	{
	// 		Title:  title,
	// 		TodoId: id,
	// 		Status: false,
	// 	},
	// })

	todo.TodoId = id
	return c.JSON(http.StatusOK, todo)
}

func getTodos(c echo.Context) error {
	d := db.New()
	defer d.Close()

	todos, err := db.GetAll(d)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "fail to fetch data!")
	}

	// return c.Render(http.StatusOK, "index", todos)
	return c.JSON(http.StatusOK, todos)
}

func getTodo(c echo.Context) error {
	idx := c.Param("id")

	id, err := strconv.Atoi(idx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id!")
	}

	d := db.New()
	defer d.Close()

	todo, err := db.GetById(d, id)
	if err != nil || todo == nil {
		// return c.Render(http.StatusBadRequest, "index", todo)
		return c.JSON(http.StatusBadRequest, todo)
	}

	// return c.Render(http.StatusOK, "index", todo)
	return c.JSON(http.StatusOK, todo)
}

func deleteTodo(c echo.Context) error {
	todo := md.Todo{}

	err := c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusForbidden, "transaction failed!")
	}

	d := db.New()
	defer d.Close()

	if err := db.Delete(d, todo.TodoId); err != nil {
		return c.JSON(http.StatusInternalServerError, "transaction failed!")
	}
	return c.JSON(http.StatusOK, "todo removed")
}

func updateTodo(c echo.Context) error {
	todo := md.Todo{}

	err := c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusForbidden, "transaction failed!")
	}

	d := db.New()
	defer d.Close()

	if err := db.Update(d, todo); err != nil {
		return c.JSON(http.StatusInternalServerError, "transaction failed!")
	}
	return c.JSON(http.StatusOK, "todo removed")
}

func Start(p string) {
	db.Init()

	t := &Template{
		template: template.Must(template.ParseFiles("public/index.html")),
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/src", "public/css")
	e.Renderer = t

	e.GET("/todos/", getTodos)
	e.GET("/todos/:id", getTodo)
	e.POST("/todos", addTodo)
	e.PUT("/todos", updateTodo)
	e.DELETE("/todos", deleteTodo)

	e.HTTPErrorHandler = ErrorPage

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", p)))
}

func ErrorPage(err error, c echo.Context) {
	code := http.StatusOK
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("public/%d.html", code)

	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
		c.JSON(http.StatusNotFound, "not found!")
	}
}

package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
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
	title := c.FormValue("title")
	if title == "" {
		return c.String(http.StatusBadRequest, "title can not be empty!")
	}

	conn := db.New()
	defer conn.Close()

	todo := md.TodoModel{
		Title:     title,
		Status:    false,
		CreatedAt: time.Now(),
	}

	id, err := db.Add(conn, todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to add todo!")
	}
	return c.Render(http.StatusOK, "index", []md.Todo{
		{
			Title:  title,
			TodoId: id,
			Status: false,
		},
	})
}

func getTodos(c echo.Context) error {
	d := db.New()
	todos, err := db.GetAll(d)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "fail to fetch data!")
	}

	return c.Render(http.StatusOK, "index", todos)
}

func getTodo(c echo.Context) error {
	idx := c.Param("id")
	if idx == "" {
		return c.String(http.StatusBadRequest, "invalid path param!")
	}

	id, err := strconv.Atoi(idx)
	if err != nil {
		return c.JSON(http.StatusForbidden, "id must be a number")
	}

	d := db.New()
	todo, err := db.GetById(d, id)
	if err != nil || todo == nil {
		return c.Render(http.StatusBadRequest, "index", todo)
	}

	return c.Render(http.StatusOK, "index", []md.Todo{})
}

func Start() {
	port := os.Getenv("PORT")

	db.Init()

	t := &Template{
		template: template.Must(template.ParseFiles("public/index.html")),
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/src", "public/css")
	e.Renderer = t

	e.GET("/todos", getTodos)
	e.GET("/todos/:id", getTodo)
	e.POST("/todos", addTodo)

	e.HTTPErrorHandler = ErrorPage

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
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
	}
}

package main

import (
	_ "fmt"
	"github.com/labstack/echo/v4"
	"github.com/ujwaldhakal/dutch-exam/service"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	lessons := os.Args[1:]

	if len(lessons) == 0 {
		log.Fatal("Please provide lessons like go run server.go 1 2 3")
	}
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		service.LoadAllQuestionAnswer(lessons...)
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/initial-question", HandleInitialQuestion)

	e.GET("/next-question", func(c echo.Context) error {
		question := c.QueryParam("question")
		answer := c.QueryParam("answer")
		service.MarkQuestionAsAsked(question, answer)
		return c.JSON(http.StatusOK, service.GetNextQuestion())
	})

	e.GET("/reset", func(c echo.Context) error {
		service.LoadAllQuestionAnswer()
		return c.JSON(http.StatusOK, service.GetNextQuestion())
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func HandleInitialQuestion(c echo.Context) error {
	return c.JSON(http.StatusOK, service.GetNextQuestion())
}

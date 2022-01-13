package main

import (
	_ "fmt"
	"github.com/labstack/echo/v4"
	"github.com/ujwaldhakal/dutch-exam/service"
	"log"
	"net/http"
	"os"
)

func main() {

	lessons := os.Args[1:]

	if len(lessons) == 0 {
		log.Fatal("Please provide lessons like go run server.go 1 2 3")
	}
	e := echo.New()

	service.LoadAllQuestionAnswer(lessons...)

	e.GET("/", func(c echo.Context) error {
		e.File("/", "public/index.html")
		return c.String(http.StatusOK, "Hello, World!")
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

package service

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"log"
	"os"
	"sync"
)

type reader func(filePath string) []byte

var readerFunc = func(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Cannot read lessons")
	}
	return data
}

type lesson struct {
	filePath string
	reader   reader
}

var totalQuestionAnswer map[string]string

func loadLesson(lesson lesson) map[string]string {

	var mutex = &sync.RWMutex{}
	mutex.Lock()
	data := lesson.reader(lesson.filePath)

	m := make(map[string]string)
	jsonError := json.Unmarshal(data, &m)

	if jsonError != nil {
		log.Fatal("Error unmarshalling json")
	}

	for key, val := range m {
		totalQuestionAnswer[key] = val
	}
	mutex.Unlock()
	return m
}

func MarkQuestionAsAsked(question string, answer string) {
	delete(totalQuestionAnswer, question)
}

func GetNextQuestion() map[string]string {
	for question, ans := range totalQuestionAnswer {
		return map[string]string{
			question: ans,
			"remaining" : fmt.Sprintf("%d", len(totalQuestionAnswer) - 1),
		}
	}
	return totalQuestionAnswer
}

func LoadAllQuestionAnswer(lessonNums ...string) map[string]string {

	totalQuestionAnswer = make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(len(lessonNums))
	for _, data := range lessonNums {
		 func(data string) {
			defer wg.Done()
			lesson := lesson{
				filePath: "service/lessons/" + data + ".json",
				reader:   readerFunc,
			}
			loadLesson(lesson)
		}(data)
	}
	wg.Wait()

	return totalQuestionAnswer
}

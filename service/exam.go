package service

import (
	"encoding/json"
	_ "fmt"
	"log"
	"os"
	"sync"
)

type reader func (filePath string) []byte

var readerFunc = func (filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Cannot read lessons")
	}
	return data
}

type lesson struct {
	filePath string
	 reader reader
}

var totalQuestionAnswer map[string]string

type askedQuestionSummary map[string]string

//func (r lesson) reader(filePath string) []byte {
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		log.Fatal("Cannot read lessons")
//	}
//	return data
//}

func loadLesson(lessonNum string) map[string]string {
	lesson := lesson{
		filePath:"service/lessons/"+lessonNum+".json",
		reader: readerFunc,
	}
	data := lesson.reader(lesson.filePath)

		m := make(map[string]string)
	jsonError := json.Unmarshal(data,&m)

	if jsonError != nil {
		log.Fatal("Error unmarshelling json")
	}

	for key,val := range m {
		totalQuestionAnswer[key] = val
	}

	return m
}

func MarkQuestionAsAsked(question string, answer string) {
	delete(totalQuestionAnswer,question)
}

func GetNextQuestion() map[string]string {
	return totalQuestionAnswer
}

func LoadAllQuestionAnswer(lessonNums ...string) map[string]string {
	totalQuestionAnswer = make(map[string]string)

	var wg sync.WaitGroup
	wg.Add(len(lessonNums))
	for _,data := range lessonNums {
		go func(data string) {
			defer wg.Done()
			loadLesson(data)
		}(data)
	}

	wg.Wait()

	return totalQuestionAnswer

}

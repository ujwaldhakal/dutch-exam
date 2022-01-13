package service

import (
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingLesson(t *testing.T) {
	totalQuestionAnswer = make(map[string]string)
	lesson := lesson{
		filePath: "asd",
		reader: func(filePath string) []byte {
			return []byte(`{"hello":"hallo"}`)
		},
	}
	data := loadLesson(lesson)
	assert.Equal(t, "hallo", data["hello"])
}

func TestLoadingMultipleLessons(t *testing.T) {
	totalQuestionAnswer = make(map[string]string)
	readerFunc = func(filePath string) []byte {
		if filePath == "service/lessons/1.json" {
			return []byte(`{"hello":"hallo"}`)
		}

		return []byte(`{"bye":"doi"}`)
	}

	data := LoadAllQuestionAnswer("1", "2")
	assert.Equal(t, 2, len(data))
}

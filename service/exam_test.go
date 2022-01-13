package service

import (
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestLoadingLesson(t *testing.T)  {
	totalQuestionAnswer = make(map[string]string)
	lesson := lesson{
		filePath: "asd",
		reader: func(filePath string) []byte {
			return []byte(`{"hello":"hallo"}`)
		},
	}
	data := loadLesson(lesson)
	assert.Equal(t,"hallo",data["hello"])
}
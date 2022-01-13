package service

import (
	"fmt"
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingLesson(t *testing.T)  {

	lesson1 := git remote add origin git@github.com:ujwaldhakal/dutch-exam.gitlesson{
		filePath: "asd",
		reader: func(filePath string) []byte {
			return []byte(`{"data":"as"}`)
		},
			}

			fmt.Println(lesson1.reader("asd"))
	assert.Equal(t,1,1)

}
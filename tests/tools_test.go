package tests

import (
	"testing"

	"github.com/cmd-e/cscalc-web/tools"
)

func TestMarksAreValid(t *testing.T) {
	testCases := []struct {
		averageMark float32
		examMark    float32
		expected    tools.ErrStruct
	}{
		{20, 50, tools.ErrStruct{IsError: false, ErrorMessage: "Информация валидна"}},
		{-20, 50, tools.ErrStruct{IsError: true, ErrorMessage: "Средний балл '-20.00' не валиден"}},
		{20, -50, tools.ErrStruct{IsError: true, ErrorMessage: "Балл за экзамен '-50.00' не валиден"}},
		{-20, -50, tools.ErrStruct{IsError: true, ErrorMessage: "Средний балл '-20.00' не валиден"}},
	}

	for i, tcase := range testCases {
		if res := tools.MarksAreValid(tcase.averageMark, tcase.examMark); res != tcase.expected {
			t.Errorf("Test number %d. Average mark: %f, exam mark: %f, expected: %v, got [%v]", i+1, tcase.averageMark, tcase.examMark, tcase.expected, res)
		}
	}
}

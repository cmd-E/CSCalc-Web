package tools

import (
	"fmt"
	"html/template"
)

// Templates stores all templates
var Templates *template.Template

// ErrStruct represent error which is sended to client
type ErrStruct struct {
	IsError      bool
	ErrorMessage string
}

// MarksAreValid checks if provided values are valid decimals above zero
func MarksAreValid(averageMark, examMark float32) ErrStruct {
	if averageMark < 0 || averageMark > 100 {
		return ErrStruct{true, fmt.Sprintf("Средний балл '%.2f' не валиден", averageMark)}
	}
	if examMark < 0 || examMark > 100 {
		return ErrStruct{true, fmt.Sprintf("Балл за экзамен '%.2f' не валиден", examMark)}
	}
	return ErrStruct{false, "Информация валидна"}
}

package calculator

// CalculateFinal calculates final mark by provided values
func CalculateFinal(averageMark, examMark float32) float32 {
	return (averageMark)*0.6 + (examMark)*0.4
}

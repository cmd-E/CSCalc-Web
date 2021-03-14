package calculator

// float finalMark = (float) (currentAverageIntFMLayout * 0.6 + examIntFMLayout * 0.4);
func CalculateFinal(averageMark int, examMark int) float32 {
	return float32(float32(averageMark)*0.6 + float32(examMark)*0.4)
}

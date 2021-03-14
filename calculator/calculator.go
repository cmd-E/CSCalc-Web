package calculator

// float finalMark = (float) (currentAverageIntFMLayout * 0.6 + examIntFMLayout * 0.4);
func CalculateFinal(averageMark int, examMark int) float64 {
	return float64(float64(averageMark)*0.6 + float64(examMark)*0.4)
}

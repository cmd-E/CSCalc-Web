package calculator

// float finalMark = (float) (currentAverageIntFMLayout * 0.6 + examIntFMLayout * 0.4);
func CalculateFinal(averageMark, examMark float32) float32 {
	return (averageMark)*0.6 + float32(examMark)*0.4
}

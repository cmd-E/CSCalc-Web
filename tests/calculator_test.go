package tests

import (
	"testing"

	"github.com/cmd-e/cscalc-web/calculator"
)

func TestCalculateFinal(t *testing.T) {
	testCases := []struct {
		averageMark float32
		examMark    float32
		expected    float32
	}{
		{11, 50, 26.6},
		{28, 12, 21.600002},
		{13.17, 53.79, 29.418001}, //13.170000, 53.799999, 29.442
		{78, 43.36, 64.144005},
		{67.53, 40.07, 56.546},
		{67.67, 38.42, 55.97},
	}

	for i, tcase := range testCases {
		if res := calculator.CalculateFinal(tcase.averageMark, tcase.examMark); res != tcase.expected {
			t.Errorf("Test number %d. Average mark: %f, exam mark: %f, expected: %f, got [%v]", i+1, tcase.averageMark, tcase.examMark, tcase.expected, res)
		}
	}
}

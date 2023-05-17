package vectorgo

import (
	"errors"
	"math"
	"testing"
)

func TestCosineSimilarity(t *testing.T) {
	testCases := []struct {
		a        []float64
		b        []float64
		expected float64
		err      error
	}{
		{[]float64{1, 1}, []float64{1, 1}, 1.0, nil},
		{[]float64{1, 1}, []float64{-1, -1}, -1.0, nil},
		{[]float64{1, 0}, []float64{0, 1}, 0.0, nil},
		{[]float64{1, 2, 3}, []float64{2, 3, 4}, 0.992583, nil},
		{[]float64{0}, []float64{}, 0.0, errors.New("vectors are not the same dimensions")},
		{[]float64{1, 2, 3}, []float64{1, 2}, 0.0, errors.New("vectors are not the same dimensions")},
		{[]float64{0, 0}, []float64{1, 1}, 0.0, errors.New("zero vector encountered")},
	}

	for _, tc := range testCases {
		result, err := CosineSimilarity(tc.a, tc.b)
		if (err != nil && tc.err == nil) || (err == nil && tc.err != nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("Expected error '%v', but got '%v'", tc.err, err)
		}

		if !floatEquals(result, tc.expected) {
			t.Errorf("CosineSimilarity(%v, %v) = %f; want %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

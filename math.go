package vectorgo

import (
	"errors"
	"math"
)

func CosineSimilarity(a []float64, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0.0, errors.New("vectors are not the same dimensions")
	}

	dotProduct := 0.0
	normA := 0.0
	normB := 0.0

	for i := 0; i < len(a); i++ {
		dotProduct += a[i] * b[i]
		normA += math.Pow(a[i], 2)
		normB += math.Pow(b[i], 2)
	}

	if normA == 0 || normB == 0 {
		return 0.0, errors.New("zero vector encountered")
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB)), nil
}

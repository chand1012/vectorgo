package db

import "gorm.io/gorm"

type Embedding struct {
	gorm.Model
	Content    string
	Identifier string `gorm:"unique"`
	Values     []EmbeddingValue
}

func (e *Embedding) ValuesToFloat64() []float64 {
	values := make([]float64, len(e.Values))
	for i, value := range e.Values {
		values[i] = value.Value
	}
	return values
}

type EmbeddingValue struct {
	gorm.Model
	Value       float64
	EmbeddingID uint
}

package vectorgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorDB(t *testing.T) {
	config := &VectorDBConfig{}
	vectorDB, err := NewVectorDB(config)
	assert.NoError(t, err)

	identifier := "test-identifier"
	content := "test content"
	embeddingData := []float64{0.1, 0.2, 0.3}

	t.Run("AddEmbedding", func(t *testing.T) {
		err := vectorDB.AddEmbedding(identifier, content, embeddingData)
		assert.NoError(t, err)
	})

	t.Run("CheckEmbeddingExists", func(t *testing.T) {
		exists, err := vectorDB.CheckEmbeddingExists(identifier)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("GetEmbeddingByIdentifier", func(t *testing.T) {
		embedding, err := vectorDB.GetEmbeddingByIdentifier(identifier)
		assert.NoError(t, err)
		assert.Equal(t, identifier, embedding.Identifier)
		assert.Equal(t, content, embedding.Content)
		assert.Equal(t, len(embeddingData), len(embedding.Values))

		for i, value := range embedding.Values {
			assert.Equal(t, embeddingData[i], value.Value)
		}
	})

	t.Run("CheckEmbeddingMatches", func(t *testing.T) {
		matches, err := vectorDB.CheckEmbeddingMatches(identifier, content)
		assert.NoError(t, err)
		assert.True(t, matches)
	})

	t.Run("DeleteEmbeddingByIdentifier", func(t *testing.T) {
		err := vectorDB.DeleteEmbeddingByIdentifier(identifier)
		assert.NoError(t, err)

		exists, err := vectorDB.CheckEmbeddingExists(identifier)
		assert.NoError(t, err)
		assert.False(t, exists)
	})
}

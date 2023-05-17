package vectorgo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmbeddingRequest(t *testing.T) {
	input := "This is a test input"
	model := "text-embedding-ada-002"
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey != "" {
		t.Run("SendValidRequest", func(t *testing.T) {
			embedding, err := sendEmbeddingRequest(input, model, apiKey)
			assert.NoError(t, err)
			assert.NotNil(t, embedding)
		})
	} else {
		t.Log("No API key provided. Skipping SendValidRequest test.")
	}

	t.Run("SendInvalidModel", func(t *testing.T) {
		invalidModel := "invalid-model"
		_, err := sendEmbeddingRequest(input, invalidModel, apiKey)
		assert.Error(t, err)
	})

	t.Run("SendInvalidApiKey", func(t *testing.T) {
		invalidApiKey := "invalid-api-key"
		_, err := sendEmbeddingRequest(input, model, invalidApiKey)
		assert.Error(t, err)
	})
}

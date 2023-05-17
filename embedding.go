package vectorgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type embeddingParams struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type embeddingObject struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     uint      `json:"index"`
}

type embeddingResponse struct {
	Embeddings []embeddingObject `json:"embeddings"`
	Object     string            `json:"object"`
	Model      string            `json:"model"`
}

func sendEmbeddingRequest(input string, model string, apiKey string) ([]float64, error) {
	params := embeddingParams{
		Model: model,
		Input: input,
	}

	jsonBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var embeddingResp embeddingResponse
	err = json.NewDecoder(resp.Body).Decode(&embeddingResp)
	if err != nil {
		return nil, err
	}

	if len(embeddingResp.Embeddings) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}

	return embeddingResp.Embeddings[0].Embedding, nil
}

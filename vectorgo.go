package vectorgo

import (
	"errors"

	"github.com/chand1012/vectorgo/db"
	"gorm.io/gorm"
)

// Config for VectorDB
type VectorDBConfig struct {
	Path      string
	OpenAIKey string
}

// VectorDB is a wrapper around a gorm.DB instance for supporting vector operations
type VectorDB struct {
	db        *gorm.DB
	openAIKey string
}

// NewVectorDB creates a new VectorDB instance
func NewVectorDB(config *VectorDBConfig) (*VectorDB, error) {
	path := config.Path

	// if path is empty, use in-memory db
	if path == "" {
		path = ":memory:"
	}

	db, err := db.NewDB(path)
	if err != nil {
		return nil, err
	}

	return &VectorDB{
		db:        db,
		openAIKey: config.OpenAIKey,
	}, nil
}

func (v *VectorDB) AddEmbedding(identifier, content string, embedding []float64) error {
	// create embedding
	embeddingModel := db.Embedding{
		Content:    content,
		Identifier: identifier,
	}

	// create embedding values
	for _, value := range embedding {
		embeddingModel.Values = append(embeddingModel.Values, db.EmbeddingValue{
			Value: value,
		})
	}

	// save embedding
	return v.db.Create(&embeddingModel).Error
}

func (v *VectorDB) GetEmbeddingByID(id uint) (db.Embedding, error) {
	var embedding db.Embedding
	err := v.db.Preload("Values").First(&embedding, id).Error
	return embedding, err
}

func (v *VectorDB) GetEmbeddingByIdentifier(identifier string) (db.Embedding, error) {
	var embedding db.Embedding
	err := v.db.Preload("Values").Where("identifier = ?", identifier).First(&embedding).Error
	return embedding, err
}

func (v *VectorDB) CheckEmbeddingExists(identifier string) (bool, error) {
	var count int64
	err := v.db.Model(&db.Embedding{}).Where("identifier = ?", identifier).Count(&count).Error
	return count > 0, err
}

// Given an identifier and content, check if the embedding's content matches
func (v *VectorDB) CheckEmbeddingMatches(identifier, content string) (bool, error) {
	var embedding db.Embedding
	err := v.db.Where("identifier = ?", identifier).First(&embedding).Error
	if err != nil {
		return false, err
	}

	return embedding.Content == content, nil
}

// Deletes the embedding. Also clears the embedding values
func (v *VectorDB) DeleteEmbeddingByID(id uint) error {
	return v.db.Where("id = ?", id).Delete(&db.Embedding{}).Error
}

func (v *VectorDB) DeleteEmbeddingByIdentifier(identifier string) error {
	return v.db.Where("identifier = ?", identifier).Delete(&db.Embedding{}).Error
}

type EmbeddingSearchResult struct {
	Embedding        db.Embedding
	CosineSimilarity float64
}

func (v *VectorDB) SearchByEmbedding(query []float64, limit uint) ([]EmbeddingSearchResult, error) {
	var embeddings []db.Embedding
	err := v.db.Preload("Values").Limit(int(limit)).Find(&embeddings).Error
	if err != nil {
		return nil, err
	}

	results := make([]EmbeddingSearchResult, len(embeddings))
	for i, embedding := range embeddings {
		cosineSimilarity, err := CosineSimilarity(query, embedding.ValuesToFloat64())
		if err != nil {
			return nil, err
		}

		results[i] = EmbeddingSearchResult{
			Embedding:        embedding,
			CosineSimilarity: cosineSimilarity,
		}
	}

	return results, nil
}

// These are the abstracted functions where OpenAI API is called

// check to make sure the openAI key is set
func (v *VectorDB) checkOpenAIKey() error {
	if v.openAIKey == "" {
		return errors.New("openAI key is not set")
	}
	return nil
}

// Adds an embedding to the database
func (v *VectorDB) Add(identifier, content string) error {
	err := v.checkOpenAIKey()
	if err != nil {
		return err
	}

	embedding, err := sendEmbeddingRequest(content, "text-embedding-ada-002", v.openAIKey)
	if err != nil {
		return err
	}

	return v.AddEmbedding(identifier, content, embedding)
}

// Searches for an embedding in plain text
func (v *VectorDB) Search(query string, limit uint) ([]EmbeddingSearchResult, error) {
	err := v.checkOpenAIKey()
	if err != nil {
		return nil, err
	}

	embedding, err := sendEmbeddingRequest(query, "text-embedding-ada-002", v.openAIKey)
	if err != nil {
		return nil, err
	}

	return v.SearchByEmbedding(embedding, limit)
}

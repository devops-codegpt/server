package models

// LLM Stores artificial intelligence LLM(large language models)
type LLM struct {
	Model
	Name string `gorm:"comment:'ai models name'" json:"name"`
}

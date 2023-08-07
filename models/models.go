package models

type Metadata struct {
	ID          string `json:"id" bson:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

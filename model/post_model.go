package model

// MPost struct
type MPost struct {
	ID          int64  `json:"id" example:"1"`
	Title       string `json:"title" example:"xyz"`
	Description string `json:"description" example:"abcdescription"`
	Status      bool   `json:"status" example:"true"`
}

// MPosts array of MPost type
type MPosts []MPost

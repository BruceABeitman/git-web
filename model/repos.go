package model

// Repo struct
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Size        int    `json:"size,omitempty"`
}

// Repos struct
type Repos []Repo

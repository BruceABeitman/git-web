package model

// SVGDetail struct
type SVGDetail struct {
	Repo     *Repo   `json:"repo,omitempty"`
	Branch   *Branch `json:"branch,omitempty"`
	Filename string  `json:"filename"`
	FileSize int64   `json:"file_size"`
	Contents string  `json:"contents"`
	Colors   []Color `json:"colors"`
}

// Color struct for color meta-data
type Color struct {
	Raw    string   `json:"raw"`
	Type   string   `json:"type"`
	Parsed []string `json:"parsed,omitempty"`
}

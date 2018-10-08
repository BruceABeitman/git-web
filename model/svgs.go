package model

// SVGList struct
type SVGList struct {
	Repo       *Repo    `json:"repo"`
	Branch     *Branch  `json:"branch"`
	TotalFiles int      `json:"total_files"`
	Filenames  []string `json:"filenames"`
}

package model

// GitBranch struct representing Git API response syntax
type GitBranch struct {
	Name   string    `json:"name"`
	Commit gitCommit `json:"commit"`
}

type gitCommit struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

// GitBranches struct representing Git API response syntax
type GitBranches []GitBranch

// Branch struct for response
type Branch struct {
	Repo   *Repo  `json:"repo,omitempty"`
	Name   string `json:"name"`
	Commit string `json:"commit"`
}

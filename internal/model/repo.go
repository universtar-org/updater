package model

type Repo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stars       int      `json:"stargazers_count"`
	Tags        []string `json:"topics"`
	Language    string   `json:"language"`
	UpdatedAt   string   `json:"updated_at"`
}

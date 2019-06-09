package model

// Paste response
type Paste struct {
	Source     string `json:"Source"`
	ID         string `json:"Id"`
	Title      string `json:"Title"`
	Date       string `json:"Date"`
	EmailCount int    `json:"EmailCount"`
}

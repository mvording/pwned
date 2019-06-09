package model

// Password response
type Password struct {
	Password string `json:"password"`
	Count    int    `json:"count"`
	Error    string `json:"error,omitempty"`
}

// PasswordResults - Array of password results
type PasswordResults []Password

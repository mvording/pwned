package model

// Breach response
type Breach struct {
	Name         string   `json:"Name"`
	Title        string   `json:"Title"`
	Domain       string   `json:"Domain"`
	BreachDate   string   `json:"BreachDate"`
	AddedDate    string   `json:"AddedDate"`
	ModifiedDate string   `json:"ModifiedDate"`
	PwnCount     int      `json:"PwnCount"`
	Description  string   `json:"Description"`
	LogoPath     string   `json:"LogoPath,omitempty"`
	DataClasses  []string `json:"DataClasses"`
	Verified     bool     `json:"IsVerified"`
	Fabricated   bool     `json:"IsFabricated"`
	Sensitive    bool     `json:"IsSensitive"`
	Retired      bool     `json:"IsRetired"`
	SpamList     bool     `json:"IsSpamList"`
}

// Breaches - set of breaches found for a given account
type Breaches struct {
	Account  string   `json:"Account"`
	Breaches []Breach `json:"Breaches"`
	Pastes   []Paste  `json:"Pastes,omitempty"`
	Error    string   `json:"Error,omitempty"`
}

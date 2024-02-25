package domain

type (
	Superhero struct {
		Name        string   `json:"name"`
		Identity    Identity `json:"identity"`
		Birthday    string   `json:"birthday"`
		Superpowers []string `json:"superpowers"`
	}
	Identity struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
)

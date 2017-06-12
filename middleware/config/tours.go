package config

type (
	Tour struct {
		Id int `json:"id"`
		Title	string	`json:"title"`
		Price	float64	`json:"price"`
		Currency	string	`json:"currency"`
		IsSpecialOffer	bool	`json:"isSpecialOffer"`
	}

	Tours struct {
		Items []Tour `json:"tours"`
	}
)

package repository

type Member struct {
	Account    string  `json:"account,omitempty"`
	Password   string  `json:"password,omitempty"`
	Permission int     `json:"permission,omitempty"`
	Name       string  `json:"name,omitempty"`
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Address    *string `json:"address,omitempty"`
	CreateAt   string  `json:"create_at,omitempty"`
}

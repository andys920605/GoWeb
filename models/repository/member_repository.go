package repository

type Member struct {
	Account    string  `json:"account,omitempty" binding:"required"`
	Password   string  `json:"password,omitempty" binding:"required"`
	Permission int     `json:"permission,omitempty" binding:"required,gte=2,lte=4"`
	Name       string  `json:"name,omitempty" binding:"required"`
	Email      *string `json:"email,omitempty" binding:"email"`
	Phone      *string `json:"phone,omitempty" binding:"len=10"`
	Address    *string `json:"address,omitempty"`
	CreateAt   string  `json:"create_at,omitempty"`
}
type UpdateMember struct {
	Account string  `json:"account,omitempty" binding:"required"`
	Name    *string `json:"name,omitempty"`
	Email   *string `json:"email,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Address *string `json:"address,omitempty"`
}

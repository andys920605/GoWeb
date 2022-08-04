package service

type Login struct {
	Account  string `json:"account,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

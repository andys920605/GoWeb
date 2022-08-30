package externals

// 寄信用
type SendMail struct {
	TargetAddress string
	Title         string
	Body          string
}

// 驗證信箱用
type VerifyEmail struct {
	Email      string `json:"email" binding:"email"`
	VerifyCode string `json:"verify_code"`
}

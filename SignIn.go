package body

// SignUpRequest Binding from JSON
type SignUpRequest struct {
	// Phone 手机号
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	// Password 密码
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// SignInRequest Binding from JSON
type SignInRequest struct {
	// 商户号
	MerchantId string `form:"merchant_id" json:"merchant_id" xml:"merchant_id"  binding:"required"`
	// Phone 手机号
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	// Password 密码
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	// Type 用户类型
	Type string `form:"type" json:"type" xml:"type"`
}
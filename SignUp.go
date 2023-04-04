package body

// 用户可以快速完成注册
// 后续补充注册信息

// SignUpRequest 快速注册
type SignUpRequest struct {
	// Phone 手机号
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	// Password 密码
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// SignUpCodeRequest 快速注册短信验证码
type SignUpCodeRequest struct {
	// Code 短信
	Code string `form:"code" json:"code" xml:"code"  binding:"required"`
}


// SignUpInfoRequest 注册信息补充
type SignUpInfoRequest struct {
	// Name 名字
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	// Email 邮箱
	Email string `form:"password" json:"password" xml:"password" binding:"required"`
}
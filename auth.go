package opt

//Auth 认证信息
type Auth interface {
	//GenerateSecret 生成密钥
	GenerateSecret() string
	//GenerateCode 生成验证码
	GenerateCode(secret string) string
	//VerifyCode 验证验证码
	VerifyCode(secret, value string) bool
}

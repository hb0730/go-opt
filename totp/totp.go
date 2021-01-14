package totp

import (
	"go-otp/hotp"
	"time"
)

//Totp
//https://github.com/xycxmz/go-totp
type Totp struct {
	// T0 in RFC6238, default 0 is OK
	// 起始时间点的时间戳
	T0 int
	//TS  in RFC6238
	// 哈希有效期的时间长度
	TS    int
	Digit int
}

func New() *Totp {
	return &Totp{
		0,
		30,
		6,
	}
}

func (t *Totp) GenerateSecret() string {
	h := &hotp.Hotp{}
	return h.GenerateSecret()
}

func (t *Totp) GenerateCode(secret string) string {
	now := (time.Now().Unix() - int64(t.T0)) / int64(t.TS)
	h := &hotp.Hotp{Counter: now, Digit: t.Digit}
	return h.GenerateCode(secret)
}
func (t *Totp) VerifyCode(secret, value string) bool {
	now := (time.Now().Unix() - int64(t.T0)) / int64(t.TS)
	h := &hotp.Hotp{Counter: now, Digit: t.Digit}
	return h.VerifyCode(secret, value)
}

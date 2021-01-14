package totp

import (
	"github.com/hb0730/go-opt"
)

//Totp
//https://github.com/xycxmz/go-totp
//https://tools.ietf.org/html/rfc6238
type Totp struct {
	opt.OTP
	interval int
}

func NewTOTP() *Totp {
	otp := opt.NewOTP(opt.GenerateSecret(), 6, nil)
	return &Totp{OTP: otp, interval: 30}
}

func (t *Totp) GenerateCode(timestamp int) string {
	return t.GenerateTOP(t.timecode(timestamp))
}

func (t *Totp) timecode(timestamp int) int {
	return int(timestamp / t.interval)
}

func (t *Totp) VerifyCode(timestamp int, value string) bool {
	return t.GenerateCode(timestamp) == value
}

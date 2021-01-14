package hotp

import (
	"github.com/hb0730/go-opt"
)

// https://tools.ietf.org/html/rfc4226
// page 5
type Hotp struct {
	opt.OTP
}

func NewHotp() *Hotp {
	otp := opt.NewOTP(opt.GenerateSecret(), 6, nil)
	return &Hotp{OTP: otp}
}

func (h *Hotp) GenerateCode(count int) string {
	return h.GenerateTOP(count)
}

func (h *Hotp) VerifyCode(count int, value string) bool {
	return h.GenerateCode(count) == value
}

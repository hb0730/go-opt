package opt

import "time"

//Totp
//https://github.com/xycxmz/go-totp
//https://tools.ietf.org/html/rfc6238
type TOTP struct {
	OTP
	//interval X rfc6238
	interval int
}

//NewTOTP create TOTP
//secret shared secret
//digits number of digits
//interval a value that reflects a time
//hasher the crypto function to use
func NewTOTP(secret string, digits, interval int, hasher *Hasher) *TOTP {
	otp := newOTP(secret, digits, hasher)
	return &TOTP{OTP: otp, interval: interval}
}

//NewDefaultTOTP create digits  6 ,interval 30, TOTP
func NewDefaultTOTP() *TOTP {
	return NewTOTP(GenerateSecret(), 6, 30, nil)
}

//GenerateCode timestamp Unix time rfc6238
func (t *TOTP) GenerateCode(timestamp int) string {
	return t.generateOTP(t.timecode(timestamp))
}

// timecode floor(timestamp / interval)
// T = (Current Unix time - T0) / X rfc6238
func (t *TOTP) timecode(timestamp int) int {
	return int(timestamp / t.interval)
}

func (t *TOTP) VerifyCode(timestamp int, value string) bool {
	return t.GenerateCode(timestamp) == value
}

//NowWithExpiration return code and expirationTime
func (t *TOTP) NowWithExpiration() (string, int64) {
	interval64 := int64(t.interval)
	timeCodeInt64 := time.Now().Unix() / interval64
	expirationTime := (timeCodeInt64 + 1) * interval64
	return t.generateOTP(int(timeCodeInt64)), expirationTime
}

package hotp

import (
	"fmt"
	"testing"
	"time"
)

func TestHotp_GenerateCode(t *testing.T) {
	now := time.Now().UTC().Unix()
	h := NewHotp()
	code := h.GenerateCode(int(now))
	fmt.Println(code)
}

func TestHotp_VerifyCode(t *testing.T) {
	now := time.Now().UTC().Unix()
	//h := &Hotp{Counter: now, Digit: 6}
	h := NewHotp()
	//secret := h.GenerateSecret()
	code := h.GenerateCode(int(now))
	b := h.VerifyCode(int(now), code)
	fmt.Println(b)
	now++
	b = h.VerifyCode(int(now), code)
	fmt.Println(b)
}

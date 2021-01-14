package hotp

import (
	"fmt"
	"testing"
	"time"
)

func TestHotp_GenerateCode(t *testing.T) {
	now := time.Now().UTC().Unix()
	h := &Hotp{Counter: now, Digit: 6}
	secret := h.GenerateSecret()
	code := h.GenerateCode(secret)
	fmt.Println(code)
}

func TestHotp_GenerateSecret(t *testing.T) {
	now := time.Now().UTC().Unix()
	h := &Hotp{Counter: now, Digit: 6}
	secret := h.GenerateSecret()
	fmt.Println(secret)
}

func TestHotp_VerifyCode(t *testing.T) {
	now := time.Now().UTC().Unix()
	h := &Hotp{Counter: now, Digit: 6}
	secret := h.GenerateSecret()
	code := h.GenerateCode(secret)
	h.Counter = now
	b := h.VerifyCode(secret, code)
	fmt.Println(b)
	b = h.VerifyCode(secret, code)
}

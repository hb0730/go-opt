package totp

import (
	"fmt"
	"testing"
	"time"
)

func TestTotp_GenerateCode(t1 *testing.T) {
	t := New()
	secret := t.GenerateSecret()
	code := t.GenerateCode(secret)
	fmt.Println(code)
}

func TestTotp_GenerateSecret(t1 *testing.T) {
	t := New()
	secret := t.GenerateSecret()
	fmt.Println(secret)
}

func TestTotp_VerifyCodeError(t1 *testing.T) {
	t := &Totp{TS: 10, Digit: 6}
	secret := t.GenerateSecret()
	code := t.GenerateCode(secret)
	fmt.Println(code)
	time.Sleep(time.Duration(20) * time.Second)
	b := t.VerifyCode(secret, code)
	fmt.Println(b)
}
func TestTotp_VerifyCode(t1 *testing.T) {
	t := &Totp{TS: 60, Digit: 6}
	secret := t.GenerateSecret()
	code := t.GenerateCode(secret)
	fmt.Println(code)
	time.Sleep(time.Duration(10) * time.Second)
	b := t.VerifyCode(secret, code)
	fmt.Println(b)
}
func TestTotp_VerifyCode2(t1 *testing.T) {
	t := &Totp{TS: 20, Digit: 6}
	secret := t.GenerateSecret()
	code := t.GenerateCode(secret)
	fmt.Println(code)
	b := t.VerifyCode(secret, code)
	fmt.Println(b)
}

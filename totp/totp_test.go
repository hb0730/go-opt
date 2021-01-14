package totp

import (
	"fmt"
	"testing"
	"time"
)

func TestTotp_GenerateCode(t1 *testing.T) {
	t := NewTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
}

func TestTotp_VerifyCode2(t1 *testing.T) {
	t := NewTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
	b := t.VerifyCode(int(now), code)
	fmt.Println(b)
}

func TestTotp_VerifyCode(t1 *testing.T) {
	t := NewTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
	time.Sleep(time.Duration(10) * time.Second)
	b := t.VerifyCode(int(now), code)
	fmt.Println(b)
}

func TestTotp_VerifyCodeError(t1 *testing.T) {
	t := NewTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
	time.Sleep(time.Duration(40) * time.Second)
	now = time.Now().UTC().Unix()
	b := t.VerifyCode(int(now), code)
	fmt.Println(b)
}

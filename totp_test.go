package opt

import (
	"fmt"
	"testing"
	"time"
)

func TestTotp_GenerateCode(t1 *testing.T) {
	t := NewDefaultTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
}

func TestTotp_VerifyCode2(t1 *testing.T) {
	t := NewDefaultTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
	b := t.VerifyCode(int(now), code)
	fmt.Println(b)
}

func TestTotp_VerifyCode(t1 *testing.T) {
	t := NewDefaultTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
	time.Sleep(time.Duration(10) * time.Second)
	b := t.VerifyCode(int(now), code)
	fmt.Println(b)
}

func TestTotp_VerifyCodeError(t1 *testing.T) {
	t := NewDefaultTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
	time.Sleep(time.Duration(40) * time.Second)
	now = time.Now().UTC().Unix()
	b := t.VerifyCode(int(now), code)
	fmt.Println(b)
}

func TestTOTP_ProvisioningUri(t1 *testing.T) {
	t := NewDefaultTOTP()
	t.secret = GenerateSecret()
	uri := t.ProvisioningUri("hb0730", "test")
	fmt.Println(uri)
}

// https://www.sooele.com/2309.html
func TestTOTP_VerifyCode(t1 *testing.T) {
	t := NewDefaultTOTP()
	t.secret = "44R4KUXMQITOA7V32M5NJJX3CZV3HGZZ"
	now := time.Now().UTC().Unix()
	b := t.VerifyCode(int(now), "185771")
	fmt.Println(b)
}
func TestTOTP_SECRET(t *testing.T) {
	fmt.Println(GenerateSecret())
}

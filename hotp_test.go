package opt

import (
	"crypto/sha512"
	"fmt"
	"testing"
	"time"
)

func TestHotp_GenerateCode(t *testing.T) {
	now := time.Now().UTC().Unix()
	h := NewDefaultHOTP()
	code := h.GenerateCode(int(now))
	fmt.Println(code)
}

func TestHotp_VerifyCode(t *testing.T) {
	now := time.Now().UTC().Unix()
	h := NewDefaultHOTP()
	code := h.GenerateCode(int(now))
	b := h.VerifyCode(int(now), code)
	fmt.Println(b)
	now++
	b = h.VerifyCode(int(now), code)
	fmt.Println(b)
}

func TestNewHOTP(t *testing.T) {
	h := NewHOTP(GenerateSecret(), 7, &Hasher{HashName: "", Digest: sha512.New})
	now := time.Now().UTC().Unix()
	code := h.generateOTP(int(now))
	b := h.VerifyCode(int(now), code)
	fmt.Println(b)
	now++
	b = h.VerifyCode(int(now), code)
	fmt.Println(b)
}

func TestHOTP_ProvisioningUri(t *testing.T) {
	now := time.Now().UTC().Unix()
	h := NewDefaultHOTP()
	h.secret = GenerateSecret()
	uri := h.ProvisioningUri("hb0730", "test", int(now))
	fmt.Println(uri)
}

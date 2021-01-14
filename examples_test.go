package opt

import (
	"fmt"
	"go-otp/totp"
	"log"
	"testing"
)

func TestGenerateSecret(t1 *testing.T) {
	t := totp.New()
	secret := t.GenerateSecret()
	log.Println(secret)
}
func TestGenerateCode(t1 *testing.T) {
	t := totp.New()
	secret := t.GenerateSecret()
	code := t.GenerateCode(secret)
	fmt.Println(code)
}
func TestVerifyCode(t1 *testing.T) {
	t := totp.New()
	secret := t.GenerateSecret()
	code := t.GenerateCode(secret)
	b := t.VerifyCode(secret, code)
	fmt.Println(b)
}

package opt

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGenerateSecret(t1 *testing.T) {
	secret := GenerateSecret()
	log.Println(secret)
}
func TestGenerateCode(t1 *testing.T) {
	t := NewDefaultTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	fmt.Println(code)
}
func TestVerifyCode(t1 *testing.T) {
	t := NewDefaultTOTP()
	now := time.Now().UTC().Unix()
	code := t.GenerateCode(int(now))
	now = time.Now().UTC().Unix()
	b := t.VerifyCode(int(now), code)
	fmt.Println(b)
}

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
	now := time.Now().UTC().Unix()
	t := newOTP(GenerateSecret(), 6, nil)
	code := t.generateOTP(int(now))
	fmt.Println(code)
}

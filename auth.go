package opt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"hash"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
)

//Auth 认证信息
type Auth interface {
	//GenerateCode 生成验证码
	GenerateCode(input int) string
	//VerifyCode 验证验证码
	VerifyCode(input int, code string) bool
}

// https://github.com/ozgur-soft/otp/blob/master/src/otp.go
type OTP struct {
	Secret  string
	Digits  int
	Content []byte
	Hasher  *Hasher
}

type Hasher struct {
	HashName string
	Digest   func() hash.Hash
}

func NewOTP(secret string, digits int, hasher *Hasher) OTP {
	if hasher == nil {
		hasher = &Hasher{
			HashName: "sha1",
			Digest:   sha1.New,
		}
	}
	return OTP{
		Secret: secret,
		Digits: digits,
		Hasher: hasher,
	}
}

func (o *OTP) byteSecret() []byte {
	missingPadding := len(o.Secret) % 8
	if missingPadding != 0 {
		o.Secret = o.Secret + strings.Repeat("=", 8-missingPadding)
	}
	sk, err := base32.StdEncoding.DecodeString(o.Secret)
	if err != nil {
		log.Println(err)
	}
	return sk
}

func (o *OTP) GenerateTOP(input int) string {
	if input < 0 {
		panic("input must be positive integer")
	}
	hasher := hmac.New(o.Hasher.Digest, o.byteSecret())
	hasher.Write(Itob(input))
	hmacHash := hasher.Sum(nil)
	offset := int(hmacHash[len(hmacHash)-1] & 0xf)
	code := ((int(hmacHash[offset]) & 0x7f) << 24) |
		((int(hmacHash[offset+1] & 0xff)) << 16) |
		((int(hmacHash[offset+2] & 0xff)) << 8) |
		(int(hmacHash[offset+3]) & 0xff)
	code = code % int(math.Pow10(o.Digits))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", o.Digits), code)
}

func GenerateSecret() string {
	s := make([]byte, 20)
	rand.Seed(time.Now().UnixNano())
	if _, err := rand.Read(s); err != nil {
		return ""
	}
	return strings.ToUpper(base32.StdEncoding.EncodeToString(s))
}
func Itob(integer int) []byte {
	byteArr := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		byteArr[i] = byte(integer & 0xff)
		integer = integer >> 8
	}
	return byteArr
}

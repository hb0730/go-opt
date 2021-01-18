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
	"net/url"
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
	secret string
	digits int
	hasher *Hasher
}

type Hasher struct {
	HashName string
	Digest   func() hash.Hash
}

const (
	OtpTypeTotp = "totp"
	OtpTypeHotp = "hotp"
)

func newOTP(secret string, digits int, hasher *Hasher) OTP {
	if hasher == nil {
		hasher = &Hasher{
			HashName: "sha1",
			Digest:   sha1.New,
		}
	}
	return OTP{
		secret: secret,
		digits: digits,
		hasher: hasher,
	}
}

func (o *OTP) byteSecret() []byte {
	missingPadding := len(o.secret) % 8
	if missingPadding != 0 {
		o.secret = o.secret + strings.Repeat("=", 8-missingPadding)
	}
	sk, err := base32.StdEncoding.DecodeString(o.secret)
	if err != nil {
		log.Println(err)
	}
	return sk
}

//OTP = HOTP(K, T)
func (o *OTP) generateOTP(input int) string {
	if input < 0 {
		panic("input must be positive integer")
	}
	hasher := hmac.New(o.hasher.Digest, o.byteSecret())
	hasher.Write(Itob(input))
	hmacHash := hasher.Sum(nil)
	offset := int(hmacHash[len(hmacHash)-1] & 0xf)
	code := ((int(hmacHash[offset]) & 0x7f) << 24) |
		((int(hmacHash[offset+1] & 0xff)) << 16) |
		((int(hmacHash[offset+2] & 0xff)) << 8) |
		(int(hmacHash[offset+3]) & 0xff)
	code = code % int(math.Pow10(o.digits))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", o.digits), code)
}

func GenerateSecret() string {
	s := make([]byte, 20)
	rand.Seed(time.Now().UnixNano())
	if _, err := rand.Read(s); err != nil {
		return ""
	}
	return strings.ToUpper(base32.StdEncoding.EncodeToString(s))
}

//BuildUri https://github.com/google/google-authenticator/wiki/Key-Uri-Format
func BuildUri(otpType, secret, accountName, issuerName, algorithm string, initialCount, digits, period int) string {
	if otpType != OtpTypeTotp && otpType != OtpTypeHotp {
		panic("otp type error, got " + otpType)
	}
	urlParams := make([]string, 0)
	urlParams = append(urlParams, "secret="+secret)
	if otpType == OtpTypeHotp {
		urlParams = append(urlParams, fmt.Sprintf("counter=%d", initialCount))
	}
	label := url.QueryEscape(accountName)
	if issuerName != "" {
		issuerNameEscape := url.QueryEscape(issuerName)
		label = issuerNameEscape + ":" + label
		urlParams = append(urlParams, "issuer="+issuerNameEscape)
	}
	if algorithm != "" && algorithm != "sha1" {
		urlParams = append(urlParams, "algorithm="+strings.ToUpper(algorithm))
	}
	if digits != 0 && digits != 6 {
		urlParams = append(urlParams, fmt.Sprintf("digits=%d", digits))
	}
	if period != 0 && period != 30 {
		urlParams = append(urlParams, fmt.Sprintf("period=%d", period))
	}
	return fmt.Sprintf("otpauth://%s/%s?%s", otpType, label, strings.Join(urlParams, "&"))
}
func Itob(integer int) []byte {
	byteArr := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		byteArr[i] = byte(integer & 0xff)
		integer = integer >> 8
	}
	return byteArr
}

package hotp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// https://tools.ietf.org/html/rfc4226
// page 5
type Hotp struct {
	//
	Counter int64
	//  digits in HOTP values
	Digit int
}

func (h *Hotp) GenerateSecret() string {
	s := make([]byte, 20)
	rand.Seed(time.Now().UnixNano())
	if _, err := rand.Read(s); err != nil {
		return ""
	}
	return strings.ToUpper(base32.StdEncoding.EncodeToString(s))
}
func (h *Hotp) GenerateCode(secret string) string {
	sk, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		log.Println(err)
	}
	text := h.toBytes(h.Counter)
	hash := h.hmacSha1(sk, text)
	code := h.truncate(hash, h.Digit)
	//使其失效
	h.Counter++
	return code
}

func (h Hotp) VerifyCode(secret, value string) bool {
	return h.GenerateCode(secret) == value
}

func (h *Hotp) toBytes(value int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, value)
	return bytesBuffer.Bytes()
}

// https://tools.ietf.org/html/rfc4226
// page 30
func (h *Hotp) hmacSha1(key, text []byte) []byte {
	client := hmac.New(sha1.New, key)
	client.Write(text)
	return client.Sum(nil)
}

// https://tools.ietf.org/html/rfc4226
// page 30
func (h *Hotp) truncate(hash []byte, digits int) string {
	length := len(hash)
	offset := hash[length-1] & 0x0f
	codeBinary :=
		(uint32(hash[offset]&0x7F) << 24) |
			(uint32(hash[offset+1]&0xFF) << 16) |
			(uint32(hash[offset+2]&0xFF) << 8) |
			(uint32(hash[offset+3] & 0xFF))
	opt := int(codeBinary) % int(math.Pow10(digits))
	result := strconv.Itoa(opt)
	for len(result) < digits {
		result = "0" + result
	}
	return result
}

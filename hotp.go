package opt

// https://tools.ietf.org/html/rfc4226
// page 5
type HOTP struct {
	OTP
}

//NewHotp  create HOTP
//secret  shared secret between client and server; each HOTP
//		  generator has a different and unique secret K.
//digits   number of digits in an HOTP value
//hasher HMAC type
func NewHOTP(secret string, digits int, hasher *Hasher) *HOTP {
	otp := newOTP(secret, digits, hasher)
	return &HOTP{OTP: otp}
}

//NewDefaultHotp  digits default 6
func NewDefaultHOTP() *HOTP {
	return NewHOTP(GenerateSecret(), 6, nil)
}

// GenerateCode
//count 8-byte counter value, the moving factor.  This counter
//           MUST be synchronized between the HOTP generator (client)
//           and the HOTP validator (server).
func (h *HOTP) GenerateCode(count int) string {
	return h.generateOTP(count)
}

//VerifyCode
//count  8-byte counter value, the moving factor
//value code
func (h *HOTP) VerifyCode(count int, value string) bool {
	return h.GenerateCode(count) == value
}

//ProvisioningUri create uri
func (h *HOTP) ProvisioningUri(accountName, issuerName string, initialCount int) string {
	return BuildUri(
		OtpTypeHotp,
		h.secret,
		accountName,
		issuerName,
		h.hasher.HashName,
		initialCount,
		h.digits,
		0)
}

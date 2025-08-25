package password

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/pbkdf2"
)

func getHash0(pass []byte) []byte {
	bs := sha256.Sum256(pass)
	return bs[:]
}

func GetHash1(password string) string {
	pass := []byte(password)
	hash0 := getHash0(pass)
	return hex.EncodeToString(hash0)
}

func getKey() []byte {
	return []byte{0x0d, 0x44, 0xa8, 0xbe, 0x3c, 0xfd, 0xab, 0x23}
}

func GetHash2(password string) string {
	hexHash1 := GetHash1(password)
	b2Hash1 := []byte(hexHash1)

	hash2 := pbkdf2.Key(b2Hash1, getKey(), 1000, 32, sha512.New)
	return hex.EncodeToString(hash2)
}

type PasswordParser struct {
}

func (p PasswordParser) Parse(data string) string {
	return GetHash3(data)
}

func GetHash3(hash1 string) string {
	b2Hash1 := []byte(hash1)

	hash2 := pbkdf2.Key(b2Hash1, getKey(), 1000, 32, sha512.New)
	return hex.EncodeToString(hash2)
}

func OtpParams() totp.ValidateOpts {
	return totp.ValidateOpts{
		Period:    43200,
		Skew:      1,
		Digits:    otp.DigitsEight,
		Algorithm: otp.AlgorithmSHA512,
	}
}

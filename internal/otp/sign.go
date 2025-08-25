package otp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/libParams"
	"github.com/hmmftg/requestCore/response"
)

type OTP struct {
	OtpKey        string
	OtpInterval   int
	OtpMinIterate int
	OtpMaxIterate int
}

func (o OTP) Hash(data string) []byte {
	hotp := hmac.New(sha256.New, []byte(o.OtpKey))
	hotp.Write([]byte(data))
	return hotp.Sum(nil)
}

func GetOTP(params libParams.ParamInterface) (*OTP, error) {
	otpIntervalString := params.GetParam("", "otpInterval")
	otpInterval, err := strconv.Atoi(*otpIntervalString)
	if err != nil {
		log.Fatalln("unable to parse otp interval", otpIntervalString, err)
		return nil, response.ToError("INVALID_OTP_PARAMS", "unable to parse otp interval", err)
	}
	otpMaxIterateString := params.GetParam("", "otpMaxIterate")
	otpMaxIterate, err := strconv.Atoi(*otpMaxIterateString)
	if err != nil {
		log.Fatalln("unable to parse otp max iterate", otpIntervalString, err)
		return nil, response.ToError("INVALID_OTP_PARAMS", "unable to parse otp interval", err)
	}
	otpMinIterateString := params.GetParam("", "otpMinIterate")
	otpMinIterate, err := strconv.Atoi(*otpMinIterateString)
	if err != nil {
		log.Fatalln("unable to parse otp min iterate", otpIntervalString, err)
		return nil, response.ToError("INVALID_OTP_PARAMS", "unable to parse otp interval", err)
	}
	otpKey := params.GetParam("", "otpKey")
	fmt.Println("otp key used", otpKey)
	return &OTP{
		OtpKey:        *otpKey,
		OtpInterval:   otpInterval,
		OtpMaxIterate: otpMaxIterate,
		OtpMinIterate: otpMinIterate,
	}, nil
	//secret := hotp.At(int(tm.Unix()) / otpInterval)
}

func (o OTP) GenOTP(terminalID string, tm time.Time) string {
	data := fmt.Sprintf("%s-%d", terminalID, int(tm.Unix())/o.OtpInterval)
	hash := base32.StdEncoding.EncodeToString(
		[]byte(
			o.Hash(data),
		))
	fmt.Println("data", data, "hash", hash)
	return hash
}

func (o OTP) checkCounter(baseCounter, minIterate, maxIterate int, terminalID string, target []byte) error {
	for i := minIterate; i < maxIterate; i++ {
		data := fmt.Sprintf("%s-%d", terminalID, baseCounter-i)
		result := o.Hash(data)
		if hmac.Equal(result, target) {
			return nil
		}
		/*log.Println("otp data", data, "invalid hash",
		base32.StdEncoding.EncodeToString(target),
		"<>",
		base32.StdEncoding.EncodeToString([]byte(result)))*/
	}
	return libError.NewWithDescription(
		http.StatusUnauthorized,
		"EXPIRED_OTP_HASH",
		"hash did't match after %d iterate", maxIterate,
	)
}

func (o OTP) CheckOTP(tm time.Time, hash, terminalID string) error {
	baseCounter := int(tm.Unix()) / o.OtpInterval
	target, errB64 := base32.StdEncoding.DecodeString(hash)
	if errB64 != nil {
		return libError.New(http.StatusUnauthorized, "BAD_OTP_HASH", errB64.Error())
	}
	return o.checkCounter(baseCounter, o.OtpMinIterate, o.OtpMaxIterate, terminalID, target)
}

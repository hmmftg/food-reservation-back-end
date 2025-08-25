package otp

import (
	"net/http"
	"time"

	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/webFramework"
)

func Check(
	HMac *OTP,
	id string,
	core requestCore.RequestCoreInterface,
	w webFramework.WebFramework,
) error {
	tm := time.Now()
	signature := HMac.GenOTP(id, tm)

	err := HMac.CheckOTP(tm, signature, id)
	if err != nil {
		return libError.New(http.StatusUnauthorized, "EXPIRED", err.Error())
	}

	tmResult, errTime := time.Parse(time.RFC3339, signature)
	if errTime != nil {
		return libError.New(http.StatusUnauthorized, "EXPIRED", "3")
	}

	tm = time.Date(2025, 8, 10, 0, 0, 0, 0, time.UTC)
	if tmResult.After(tm) {
		return libError.New(http.StatusUnauthorized, "EXPIRED", "4")
	}

	return nil
}

package ums

import (
	"encoding/base32"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hmmftg/food-reservation-back-end/internal/password"
	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/webFramework"
	"github.com/pquerna/otp/totp"
)

func GenJwtKey() jwt.Keyfunc {
	return func(t *jwt.Token) (any, error) {
		return []byte(GetOTPKey()), nil
	}
}

func GetOTPKey() string {
	return "50cw7IEJqo*VT3JF7x5Y^9A6^tdBY07g"
}

func GenerateOtp(key string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(GetOTPKey() + key))
	options := password.OtpParams()
	passCode, err := totp.GenerateCodeCustom(secret, time.Now(), options)
	if err != nil {
		panic(err)
	}
	return passCode
}

func GetToken(w webFramework.WebFramework) (string, error) {
	authHeader := w.Parser.GetHeaderValue("Authorization")
	if len(authHeader) == 0 {
		return "", libError.New(http.StatusBadRequest,
			"AUTH_HEADER_ABSENT_OR_INVALID", "auth header does not exists")
	}
	auth := strings.SplitN(authHeader, " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		return "", libError.New(http.StatusUnauthorized, "AUTH_BAD_METHOD",
			"auth header type is not valid")
	}
	return auth[1], nil

}

func GenJwtToken(payload jwt.RegisteredClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(GetOTPKey()))
}

func GenerateToken(
	dt time.Time,
	ageNum int,
	subject string,
	audience []string,
	id string,
) (string, error) {
	dtValidUntil := dt.Add(time.Second * time.Duration(ageNum))
	//token := BearerToken{User: user, Flags: flagsRaw, Roles: roles, Start: dt, ValidUntil: dtValidUntil.Format("20060102150405")}
	otpSecret := GenerateOtp(id + dt.Format("20060102150405"))
	claims := jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: dtValidUntil},
		NotBefore: &jwt.NumericDate{Time: dt},
		IssuedAt:  &jwt.NumericDate{Time: dt},
		Subject:   subject,
		Audience:  audience,
		ID:        id,
		Issuer:    otpSecret,
	}
	jwtToken, err := GenJwtToken(claims)
	if err != nil {
		log.Printf("GenJwt()=>Payload: %+v, Error: %+v\n", claims, err)
		return "AUTH_JWT_SIGN_ERROR", errors.Join(err, errors.New("Error Signing Token"))
	}
	return jwtToken, nil
}

func ValidateOtp(key, passCode string) bool {
	secret := base32.StdEncoding.EncodeToString([]byte(GetOTPKey() + key))
	options := password.OtpParams()
	isValid, err := totp.ValidateCustom(passCode, secret, time.Now(), options)
	if err != nil {
		panic(err)
	}
	return isValid
}

func ValidateJwtToken(
	core requestCore.RequestCoreInterface,
	tokenRaw string,
) (*User, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenRaw, &jwt.RegisteredClaims{}, GenJwtKey())
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && validationErr.Is(jwt.ErrTokenExpired) {
			return nil, libError.New(http.StatusUnauthorized, "EXPIRED_TOKEN", "expired token")
		}

		return nil, libError.New(http.StatusUnauthorized, "ERROR_PARSE_TOKEN", err.Error())
	}
	token := jwtToken.Claims.(*jwt.RegisteredClaims)
	dtStart := time.Unix(token.IssuedAt.Unix(), 0).UTC()
	dtEnd := time.Unix(token.ExpiresAt.Unix(), 0).UTC()
	dtNow := time.Now().UTC()
	log.Println("Diff:", dtNow.Sub(dtStart), dtEnd.Sub(dtNow), token.Issuer)
	if !(dtNow.After(dtStart) && dtNow.Before(dtEnd)) {
		return nil, libError.New(http.StatusUnauthorized, "EXPIRED_TOKEN", "expired token")
	}

	if !ValidateOtp(token.ID+dtStart.Format("20060102150405"), token.Issuer) {
		return nil, libError.New(http.StatusUnauthorized, "INVALID_TOKEN", "invalid token")
	}

	usr, getUserErr := GetUserData(token.ID, core)
	if getUserErr != nil {
		return nil, libError.New(http.StatusUnauthorized, "USER_NOT_FOUND", getUserErr.Error())
	}
	core.Params().GetRemoteApi("locker")

	return usr, nil
}

package ums

import (
	"github.com/hmmftg/food-reservation-back-end/internal/otp"
	"github.com/hmmftg/food-reservation-back-end/internal/params"
	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libParams"
)

type umsEnv struct {
	Params    libParams.ParamInterface
	Interface requestCore.RequestCoreInterface
}

type CheckRequest struct {
}

type CheckResponse struct {
	UserId        string   `json:"id"`
	UserName      string   `json:"name"`
	PersonID      string   `json:"person"`
	Authenticated bool     `json:"authenticated"`
	Roles         []string `json:"roles"`
	Flags         []string `json:"flags"`
}

type CheckHandler struct {
	Name   string
	UserID string
	HMac   *otp.OTP
}

type AuthHeader struct {
	Authentication string `header:"Authorization" validate:"startswith=Bearer"`
}

type ServiceAuthHandler struct {
}

type AuthHandlerInterface interface {
	// main handler runs after initialize
	Handler(core requestCore.RequestCoreInterface, req AuthHeader) error
}

type LoginRequest struct {
	UserName string `json:"username"`
	Pass     string `json:"password"`
}

type Role struct {
	params.Model
}

type User struct {
	params.Model
	Name       string
	Department string
	Roles      []Role `gorm:"many2many:user_roles;"`
	PersonID   string
	Data       string
	Password   string
}

type LoginHandler struct {
	Name string
	HMac *otp.OTP
}

type ValidationError struct {
	Field string
	Error string
}

type HttpError struct {
	Status  int
	Message string
	Errors  []ValidationError
}

type LoginResponse struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	PersonID     string     `json:"personID"`
	Department   string     `json:"department"`
	Roles        []string   `json:"roles"`
	UserData     string     `json:"userData"`
	RefreshToken string     `json:"refresh_token"`
	AccessToken  string     `json:"access_token"`
	Error        *HttpError `json:"error"`
}

type LogoutRequest struct {
	ID string `json:"id"`
}

type LogoutResponse struct {
	State string `json:"state"`
}

type LogoutHandler struct {
	Name string
}

type RegisterRequest struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Pass     string `json:"password"`
}

type RegisterResponse struct {
	Error string `json:"error"`
}

type RegisterHandler struct {
	Name string
}

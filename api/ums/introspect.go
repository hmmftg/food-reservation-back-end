package ums

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libContext"
	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/libValidate"
	"github.com/hmmftg/requestCore/response"
)

func ParseHeader(header *AuthHeader) error {
	err, errValidate := libValidate.ValidateStruct(header)
	if err != nil {
		return libError.New(http.StatusInternalServerError, "INVALID_VALIDATION", err)
	}
	if errValidate != nil {
		return libError.New(http.StatusUnauthorized, "INVALID_AUTH_HEADER", errValidate)
	}
	return nil
}

func (a ServiceAuthHandler) Handler(core requestCore.RequestCoreInterface, req AuthHeader) error {
	parts := strings.Split(req.Authentication, " ")
	if len(parts) != 2 {
		return libError.New(
			http.StatusUnauthorized,
			"AUTH_HEADER_ABSENT_OR_INVALID",
			"auth header has invalid format",
		)
	}
	_, err := ValidateJwtToken(core, parts[1])
	if err != nil {
		return err
	}

	return nil
}

func (env umsEnv) UmsIntrospect(title string, handler AuthHandlerInterface) any {
	log.Println("Registering: auth middleware")
	return func(c context.Context) {
		w := libContext.InitContextNoAuditTrail(c)
		core := env.Interface
		defer func() {
			if r := recover(); r != nil {
				switch data := r.(type) {
				case error:
					core.Responder().Error(w,
						libError.NewWithDescription(
							http.StatusInternalServerError,
							response.SYSTEM_FAULT,
							"error in %s", title,
						))
				default:
					core.Responder().Error(w,
						libError.NewWithDescription(
							http.StatusInternalServerError,
							response.SYSTEM_FAULT,
							"error in %s=> %+v", title, data,
						))
				}
				errAbort := w.Parser.Abort()
				if errAbort != nil {
					log.Println("error abort", errAbort)
				}
				panic(r)
			}
		}()

		header := AuthHeader{
			Authentication: w.Parser.GetHeaderValue("Authorization"),
		}
		err := ParseHeader(&header)
		if err != nil {
			core.Responder().Error(w, err)
			errAbort := w.Parser.Abort()
			if errAbort != nil {
				log.Println("error abort", errAbort)
			}
			return
		}

		err = handler.Handler(core, header)
		if err != nil {
			core.Responder().Error(w, err)
			errAbort := w.Parser.Abort()
			if errAbort != nil {
				log.Println("error abort", errAbort)
			}
			return
		}

		errNext := w.Parser.Next()
		if errNext != nil {
			log.Println("error next", errNext)
		}
	}
}

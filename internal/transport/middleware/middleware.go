package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	jwt2 "github.com/MediStatTech/MediStat-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const headerPersonalId = "Personal-Id"

func RecoverMiddleware() echo.MiddlewareFunc {
	config := middleware.DefaultRecoverConfig
	config.LogErrorFunc = func(c echo.Context, err error, stack []byte) error {
		zap.L().Sugar().Errorf("PANIC RECOVER: %v %s", err, stack)
		return err
	}

	return middleware.RecoverWithConfig(config)
}

func CORSMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8443"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	})
}

func JWT(verifier jwt2.JWT) func(next echo.HandlerFunc) echo.HandlerFunc {
	return echojwt.WithConfig(echojwt.Config{
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := verifier.Verify(auth)
			if err != nil {
				return nil, err
			}
			c.Request().Header.Set(headerPersonalId, claims.Subject)
			return claims, nil
		},
	})
}

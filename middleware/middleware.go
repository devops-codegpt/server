package middleware

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/request"
	"github.com/golang-jwt/jwt/v4"
	casbinMw "github.com/labstack/echo-contrib/casbin"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"regexp"
)

func Init(e *echo.Echo, c container.Container) {
	e.Use(RecoverMiddleware())
	e.Use(RequestLogMiddleware(c))
	e.Use(JwtMiddleware(c))
	e.Use(CasbinAuthMiddleware(c))
}

func RecoverMiddleware() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
	})
}

func CorsMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
		AllowOrigins:                             []string{"*"},
		AllowHeaders: []string{
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderAcceptEncoding,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		MaxAge: 86400,
	})
}

func RequestLogMiddleware(cr container.Container) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger := cr.GetLogger().GetZapLogger()
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	})
}

func JwtMiddleware(cr container.Container) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper:    skipper,
		SigningKey: []byte(cr.GetConfig().Auth.JwtKey),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &request.JwtCustomClaims{}
		},
	})
}

func CasbinAuthMiddleware(cr container.Container) echo.MiddlewareFunc {
	return casbinMw.MiddlewareWithConfig(casbinMw.Config{
		Skipper:  skipper,
		Enforcer: cr.GetAuthentication().GetCasbinEnforcer(),
		// The role is used for authentication here, so UserGetter needs to be rewritten
		UserGetter: func(c echo.Context) (string, error) {
			// Get role information from jwt
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*request.JwtCustomClaims)
			roleKeyword := claims.RoleKeyword
			return roleKeyword, nil
		},
		ErrorHandler: nil,
	})
}

var skipper middleware.Skipper = func(c echo.Context) bool {
	return equalPath(
		c.Path(),
		[]string{
			config.APIHealth,
			config.APIUserLogin,
			config.APIConversationWS,
			config.APIConversationBase,
			config.APICodeGenerate,
		},
	)
}

// equalPath judges whether a given path contains in the path list.
func equalPath(cpath string, paths []string) bool {
	for i := range paths {
		if regexp.MustCompile(paths[i]).MatchString(cpath) {
			return true
		}
	}
	return false
}

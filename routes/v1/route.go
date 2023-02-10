package v1

import (
	// "discusiin/controllers/topics"

	"io"
	"net/http"
	"projek_fisioterapi/configs"
	"projek_fisioterapi/controllers/users"
	mid "projek_fisioterapi/middleware"
	"projek_fisioterapi/routes"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(payload *routes.Payload) (*echo.Echo, io.Closer) {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	mid.LogMiddleware(e)
	e.Use(middleware.Recover())
	cors := middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPut,
				http.MethodPatch,
				http.MethodPost,
				http.MethodDelete},
			AllowHeaders: []string{
				"Accept",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"X-CSRF-Token",
				"Authorization",
				"Origin",
			},
		})
	e.Use(cors)

	trace := jaegertracing.New(e, nil)

	uHandler := users.UserHandler{
		IUserServices: payload.GetUserServices(),
	}

	api := e.Group("/api")
	v1 := api.Group("/v1")

	//endpoints users
	users := v1.Group("/users")
	users.GET("", uHandler.GetProfile, middleware.JWT([]byte(configs.Cfg.TokenSecret)))

	return e, trace
}

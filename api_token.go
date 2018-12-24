package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	captcha "github.com/pokkystr/golang-api/Captcha"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello/go", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello GoLang \n")
	})

	e.GET("/hello/:name", func(c echo.Context) error {
		return c.String(http.StatusOK, "/hello/"+c.Param("name"))
	})

	e.POST("/user/info", func(c echo.Context) error {
		var userName = "pigke"
		var password = "kulk"

		return c.String(http.StatusOK, "userName: "+userName+", password: "+password)
	})

	e.POST("/token/gen", func(c echo.Context) error {
		var captcha = captcha.GenCaptcha(1, 1, 1, 3)
		return c.String(http.StatusOK, captcha)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

package handler

import (
	c "github.com/blue-farid/WebMidExam/config"
	"github.com/blue-farid/WebMidExam/security"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST(c.Cfg.Route.Signup, SignupHandler)
	e.POST(c.Cfg.Route.Login, LoginHandler)
	e.GET(c.Cfg.Route.Basket+"/:id", GetBasketHandler, security.JWTMiddleware)
	e.POST(c.Cfg.Route.Basket, CreateBasketHandler, security.JWTMiddleware)
	e.PATCH(c.Cfg.Route.Basket+"/:id", UpdateBasketHandler, security.JWTMiddleware)
	e.GET(c.Cfg.Route.Basket, GetBasketHandler, security.JWTMiddleware)
	e.DELETE(c.Cfg.Route.Basket+"/:id", DeleteHandler, security.JWTMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}

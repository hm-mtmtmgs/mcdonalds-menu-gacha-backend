package router

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/env"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/router/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(env env.Env) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(env.LoggerLevel)
	e.Validator = middlewares.NewValidatorMiddleware()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	return e
}

func Register(r *echo.Echo, controllerManager *controllers.Manager) {
	jwtAuth := middlewares.NewJwtMiddleware()

	rv := r.Group("/v1")
	rv.POST("/signup", controllerManager.UserController.SignUp)
	rv.POST("/login", controllerManager.UserController.Login)
	rv.GET("/user", controllerManager.UserController.GetUser, jwtAuth)
	rv.GET("/menus", controllerManager.MenuController.GetMenuList, jwtAuth)
	rv.GET("/menu-gacha", controllerManager.MenuController.GetMenuGacha)
}

package main

import (
	"github.com/flowkater/go-ddd-sample/src/config"
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"

	interfaces "github.com/flowkater/go-ddd-sample/src/interfaces/todo"
	"github.com/flowkater/go-ddd-sample/src/module"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	if err := config.InitDB(); err != nil {
		panic(err.Error())
	}

	db := config.DB()
	if err := db.AutoMigrate(&user_domain.User{}, &todo_domain.Todo{}); err != nil {
		panic(err.Error())
	}

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	// e.Use(config.ContextLogger())
	e.Use(middleware.Gzip())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{os.Getenv("CORS_ALLOW_ORIGIN")},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	todoApiController := module.InitializeTodoApiController()

	interfaces.RegiterHandlerTodos(e, todoApiController)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":8000"))
}

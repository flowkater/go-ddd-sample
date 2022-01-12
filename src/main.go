package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/flowkater/go-ddd-sample/src/config"
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"github.com/flowkater/go-ddd-sample/src/interfaces/resolver"
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

	registerGraphqlHandler(e)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":8000"))
}

func registerGraphqlHandler(e *echo.Echo) {
	todoFacade := module.InitializeTodoFacade()

	config := resolver.Config{
		Resolvers: resolver.NewResolver(todoFacade),
	}

	graphqlHandler := func(c echo.Context) error {
		return GraphqlHandler(c, config)
	}

	e.POST("/graphql", graphqlHandler)
}

func GraphqlHandler(c echo.Context, config resolver.Config) error {
	h := handler.NewDefaultServer(resolver.NewExecutableSchema(config))

	r := c.Request()

	// if c.Get("user") != nil {
	// 	user := auth.GetUserByJWT(c)
	// 	if user != nil {
	// 		ctx := context.WithValue(r.Context(), auth.UserCtxKey, user)
	// 		r = r.WithContext(ctx)

	// 		ctx = context.WithValue(r.Context(), auth.IpCtxKey, c.RealIP())
	// 		r = r.WithContext(ctx)

	// 		// ctx = context.WithValue(r.Context(), auth.DateChangeTimeCtxKey, r.Header.Get("Date-Change-Time"))
	// 		// r = r.WithContext(ctx)
	// 	}
	// }

	h.ServeHTTP(c.Response(), r)

	// .SetRecoverFunc(func(ctx context.Context, err interface{}) error {
	// 	// notify bug tracker...

	// 	return errors.New("Internal server error!")
	// })

	return nil
}

// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }

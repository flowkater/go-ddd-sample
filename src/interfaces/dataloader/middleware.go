package dataloader

import (
	"context"
	"net/http"

	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"github.com/labstack/echo/v4"
)

var loadersKey = &contextKey{"dataloaders"}

type contextKey struct {
	name string
}

type Loaders struct {
	UserById user_domain.UserLoader
}

type Facades struct {
	TodoFacade application.TodoFacade
	UserFacade application.UserFacade
}

func DataloaderMiddleware(e echo.Context, next http.Handler, facades Facades) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := &Loaders{
			UserById: facades.UserFacade.UserByIdLoader(e.Request().Context()),
		}
		ctx := context.WithValue(r.Context(), loadersKey, loaders)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

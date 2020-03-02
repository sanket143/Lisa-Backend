package router

import (
	"net/http"

	"github.com/sanket143/lisa-backend/pkg/types/routes"
	"github.com/sanket143/lisa-backend/src/router/caves"
)

// Middleware function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes function
func GetRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{"Home", "GET", "/", caves.HomeHandler},
	}
}

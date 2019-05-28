package appRouter

import "net/http"

type AppRoute struct {
	Name    string
	URL     string
	Handler http.HandlerFunc
	Method  string
}

type AppRoutes []AppRoute

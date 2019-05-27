package appRouter

import "net/http"

type AppRoute struct {
	Name    string
	URL     string
	Handler http.HandlerFunc
	Method  string
}

type AppRoutes []AppRoute

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Qunatity int    `json:"qty,omitempty"`
}

var products []Product

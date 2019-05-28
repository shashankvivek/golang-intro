package appRouter

import (
	app "myApp/products"

	"github.com/gorilla/mux"
)

var routeList = AppRoutes{

	AppRoute{
		"GetProducts",
		"/products",
		app.GetProducts,
		"GET",
	},
	AppRoute{
		"GetProductFromID",
		"/product/{id}",
		app.GetProductById,
		"GET",
	},
	AppRoute{
		"CreateProduct",
		"/product",
		app.CreateProduct,
		"POST",
	},
	AppRoute{
		"DELETEProduct",
		"/product",
		app.DeleteProduct,
		"DELETE",
	},
}

func NewRouter() *mux.Router {
	muxRouter := mux.NewRouter()
	for _, route := range routeList {
		muxRouter.
			Methods(route.Method).
			Path(route.URL).
			Name(route.Name).
			Handler(route.Handler)
	}
	return muxRouter
}

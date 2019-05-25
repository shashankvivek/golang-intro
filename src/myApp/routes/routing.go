package appRouter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	products = append(products, Product{ID: "id1", Name: "Prod1", Qunatity: 2})
	products = append(products, Product{ID: "id2", Name: "Prod2", Qunatity: 5})
	muxRouter := mux.NewRouter()
	for _, route := range routeList {
		muxRouter.Methods(route.Method).Path(route.URL).Name(route.Name).Handler(route.Handler)
	}
	return muxRouter
}

func GetProducts(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(products)
}

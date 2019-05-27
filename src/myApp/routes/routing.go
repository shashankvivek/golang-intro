package appRouter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var routeList = AppRoutes{
	AppRoute{
		"GetProducts",
		"/products",
		GetProducts,
		"GET",
	},
	AppRoute{
		"GetProductFromID",
		"/product/{id}",
		GetProductById,
		"GET",
	},
	AppRoute{
		"CreateProduct",
		"/product",
		CreateProduct,
		"POST",
	},
	AppRoute{
		"DELETEProduct",
		"/product",
		DeleteProduct,
		"DELETE",
	},
}

func NewRouter() *mux.Router {
	products = append(products, Product{ID: "id1", Name: "Prod1", Qunatity: 2})
	products = append(products, Product{ID: "id2", Name: "Prod2", Qunatity: 5})
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

func GetProducts(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func GetProductById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	var prod Product
	_ = json.NewDecoder(req.Body).Decode(&prod)
	// d, _ := json.Marshal(&prod)
	fmt.Println(prod)
	products = append(products, prod)
	json.NewEncoder(w).Encode(&products)
}

func DeleteProduct(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}

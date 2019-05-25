package main

import (
	"log"
	"net/http"

	app "myApp/routes"
)

// type Product struct {
// 	ID       string `json:"id"`
// 	Name     string `json:"name"`
// 	Qunatity int    `json:"qty,omitempty"`
// }

// var products []Product

func main() {
	handleRequest()
}

func handleRequest() {
	// products = append(products, Product{ID: "id1", Name: "Prod1", Qunatity: 2})
	// products = append(products, Product{ID: "id2", Name: "Prod2", Qunatity: 5})
	// r := mux.NewRouter()
	// r.HandleFunc("/products", GetProducts).Methods("GET")
	// r.HandleFunc("/product/{id}", GetProductById).Methods("GET")
	// r.HandleFunc("/product", CreateProduct).Methods("POST")
	// r.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
	r := app.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

// func GetProducts(w http.ResponseWriter, req *http.Request) {
// 	json.NewEncoder(w).Encode(products)
// }

// func GetProductById(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for _, item := range products {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Product{})
// }

// func CreateProduct(w http.ResponseWriter, req *http.Request) {
// 	var prod Product
// 	_ = json.NewDecoder(req.Body).Decode(&prod)
// 	// d, _ := json.Marshal(&prod)
// 	fmt.Println(prod)
// 	products = append(products, prod)
// 	json.NewEncoder(w).Encode(&products)
// }

// func DeleteProduct(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for index, item := range products {
// 		if item.ID == params["id"] {
// 			products = append(products[:index], products[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(products)
// }

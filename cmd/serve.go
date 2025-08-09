package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

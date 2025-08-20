package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"log"
	"net/http"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	log.Println("AMI Data dibo")
}

func Server() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	// wrappedMux := manager.WrapMux(
	// 	mux,
	// 	middleware.Logger,
	// 	middleware.Hudai,
	// 	middleware.CorsWithPreflight,
	// )

	globalMiddlewares := []middleware.Middleware{
		middleware.CorsWithPreflight,
		middleware.Hudai,
		middleware.Logger,
	}

	manager.Use(middleware.CorsWithPreflight)

	wrappedMux := manager.WrapMux(globalMiddlewares, mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

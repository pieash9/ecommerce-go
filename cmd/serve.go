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
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

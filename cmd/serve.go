package cmd

import (
	"ecommerce/global_router"
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

	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

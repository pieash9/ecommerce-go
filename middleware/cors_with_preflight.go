package middleware

import "net/http"

func CorsWithPreflight(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle cors issues
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Client-Id")
		w.Header().Set("Content-Type", "Application/json")

		// handle preflight requests
		w.WriteHeader(http.StatusOK)
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

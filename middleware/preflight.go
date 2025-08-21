package middleware

import "net/http"

func Preflight(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle preflight requests
		w.WriteHeader(http.StatusOK)
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

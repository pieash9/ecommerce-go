package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Client-Id")
		w.Header().Set("Content-Type", "Application/json")

		w.WriteHeader(http.StatusOK)
		if r.Method == http.MethodOptions {
			return
		}

		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllReq)

}

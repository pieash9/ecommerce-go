package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middleware ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	// for _, globalMiddleware := range mngr.globalMiddlewares {
	// 	h = globalMiddleware(h)
	// }

	return h
}

func (mngr *Manager) WrapMux(middlewares []Middleware, handler http.Handler) http.Handler {
	h := handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

package middlewares

import (
	"net/http"
)

type Middleware struct{}

func (m *Middleware) routerWrapper(w http.ResponseWriter, r *http.Request,
	router http.HandlerFunc) {
	if r.URL.Path == "/" {
		w.Write([]byte("Hello world!"))
	}
}

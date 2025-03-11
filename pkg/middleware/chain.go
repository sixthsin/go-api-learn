package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(nextHandler http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			nextHandler = middlewares[i](nextHandler)
		}
		return nextHandler
	}
}

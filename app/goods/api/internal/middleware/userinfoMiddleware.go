package middleware

import "net/http"

type UserInfoMiddleware struct {
}

func NewUserInfoMiddleware() *UserInfoMiddleware {
	return &UserInfoMiddleware{}
}

func (m *UserInfoMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}

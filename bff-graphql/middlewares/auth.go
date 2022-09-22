package middlewares

import (
	"net/http"

	"google.golang.org/grpc/metadata"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			ctx := r.Context()
			if len(token) != 0 {
				ctx = metadata.NewIncomingContext(r.Context(), metadata.New(map[string]string{"authorization": token}))
			}

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

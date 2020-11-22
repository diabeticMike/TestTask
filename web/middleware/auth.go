package middleware

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/TestTask/logger"
	"github.com/TestTask/repository"
)

// NewAuthMiddleware create AuthMiddleware instance
func NewAuthMiddleware(log logger.Logger, authRepo repository.AuthRepository) *AuthMiddleware {
	return &AuthMiddleware{log: log, authRepo: authRepo}
}

// AuthMiddleware struct for holding additional info about authorization
type AuthMiddleware struct {
	log      logger.Logger
	authRepo repository.AuthRepository
}

// AuthMiddleware is jwt authorization middleware
func (am *AuthMiddleware) AuthMiddleware(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Api-key")
		log.Println(r.URL.Path)

		if key == "" {
			am.log.Errorf("empty api-key header error in %s\n", r.URL.Path)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(`Missing api-key`))
			return
		}

		if _, err := am.authRepo.GetAuthByAPIKey(key); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				am.log.Errorf("Wrong api-key\n")
				w.WriteHeader(http.StatusForbidden)
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`Missing api-key`))
				return
			}
			am.log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(err.Error()))
			return
		}

		am.log.Printf("%v key used for authorization", key)
		next.ServeHTTP(w, r)
	})
}

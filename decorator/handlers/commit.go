package handlers

import (
	"net/http"
)

func cacheHash() string {
	return "TODO"
}

// WithHash starts OMIT
func WithHash(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-CACHE-HASH", cacheHash())
		h(w, r)
	}
}

// END OMIT

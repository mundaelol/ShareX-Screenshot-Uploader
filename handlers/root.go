package handlers

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "You're not supposed to be here", http.StatusForbidden)
}

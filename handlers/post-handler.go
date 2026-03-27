package handlers

import (
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a Post Response"))
}

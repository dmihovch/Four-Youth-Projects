package network

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.Ser)
	}
}

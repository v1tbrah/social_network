package internal

import (
	"io"
	"net/http"
)

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "There is a connection.")
}

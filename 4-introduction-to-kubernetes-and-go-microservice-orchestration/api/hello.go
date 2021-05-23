package api

import (
	"fmt"
	"net/http"
)

func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go")
}

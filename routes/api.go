package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	_ = http.ListenAndServe(":9990", route)
}

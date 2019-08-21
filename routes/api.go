package routes

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func Init() {
  route := mux.NewRouter()

  route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

  http.ListenAndServe(":9990", route)
}

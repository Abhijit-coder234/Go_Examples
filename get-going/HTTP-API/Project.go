package main 

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong"
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)
}
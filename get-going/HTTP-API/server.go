package main

import "net/http"

func main() {
	mux := http.NewServeMux() //handler for http
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("request received")
		fmt.Println(r.Method)  // to return the type of the request
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe("localhost:3000", mux)
}

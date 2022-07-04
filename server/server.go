package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, req *http.Request) {
		message := req.URL.Query().Get("message")
		fmt.Fprintf(w, fmt.Sprintf("<h1>%s</h1>", message))

	})
	http.ListenAndServe(":8081", nil)

}

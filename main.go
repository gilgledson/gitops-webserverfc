package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hellow full cycle 03</h1>"))
	})
	http.ListenAndServe(":8001", nil)
}

package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	fileServerMux := http.NewServeMux()
	fileServerMux.Handle("/", fileServer)
	http.ListenAndServe(":8080", fileServerMux)
}

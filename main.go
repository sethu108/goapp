package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, AWS CodePipeline!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8090...")
	http.ListenAndServe(":8090", nil)
}

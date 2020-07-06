package main

import (
	"fmt"
	"net/http"

	"pointer_demonstrate/nonpointer"
	"pointer_demonstrate/pointer"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/pointer_demonstrate", pointer.New().Demonstrate)
	http.HandleFunc("/nonpointer_demonstrate", nonpointer.New().Demonstrate)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Root")
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

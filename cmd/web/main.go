package main

import (
	"fmt"
	"net/http"

	"github.com/yanuarultfah/golang/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	// fmt.Println("Hello world")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello world")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Printf("Number Bytes written:%d", n))
	// })
	// solution 2
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting application on port%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

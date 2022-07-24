package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is homepage")
	renderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is About Page and 2+2 is %d", sum))
	renderTemplate(w, "about.page.tmpl")
}
func addValues(x, y int) int {
	return x + y
}
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "can not divided by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Can not divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

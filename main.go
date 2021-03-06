package main

import (
	"fmt"
	"net/http"
)

func main() {
	Print("git learning")
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func Add(a ...int) (sum int) {
	for _, i := range a {
		sum += i
	}
	return sum
}

func Sub(a ...int) (sub int) {
	for _, i := range a[1:] {
		sub -= i
	}
	return a[0] + sub
}

func Handler(request http.Request) http.Response {
	return http.Response{}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

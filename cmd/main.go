package main

import (
	"net/http"
	"service/internal/handlers"
	"service/internal/stan"
)

func main() {
	http.HandleFunc("/order", handlers.HandlerOrder)

	stan.Stan()

	http.ListenAndServe(":8080", nil)
}

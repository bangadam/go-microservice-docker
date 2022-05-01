package main

import (
	"log"
	"net/http"

	"github.com/bangadam/go-microservice-docker/handlers"
	"github.com/bangadam/go-microservice-docker/models"
	"github.com/bangadam/go-microservice-docker/repositories"
)

func main() {
	// repository
	repositories.AddProduct(models.Product{
		Name: "Product 1",
		Price: 100,
	})

	repositories.AddProduct(models.Product{
		Name: "Product 2",
		Price: 200,
	})

	// handler
	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
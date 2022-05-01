package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bangadam/go-microservice-docker/models"
	"github.com/bangadam/go-microservice-docker/repositories"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Received: ", r.Method)

	switch r.Method {
		case http.MethodGet:
			list(w, r)
			break
		case http.MethodPost:
			add(w, r)
			break
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method Not Allowed"))
			break
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	products := repositories.GetProducts()
	json, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

	log.Println("Response Sent: ", http.StatusOK)
}

func add(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)

	var product models.Product
	err := json.Unmarshal(payload, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	product.ID = repositories.AddProduct(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	json, _ := json.Marshal(product)
	w.Write(json)

	log.Println("Response Sent: ", http.StatusCreated)
}
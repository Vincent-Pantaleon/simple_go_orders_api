package handlers

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an Order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Orders")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Order by ID")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an Order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}

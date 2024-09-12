package main

import (
	"encoding/json"
	"net/http"
)

// Definimos los structs para las piezas y productos
type Pieza struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

// Datos iniciales
var piezas = []Pieza{
	{ID: 1, Nombre: "Pieza 1", Descripcion: "Suite con jacuzzi"},
	{ID: 2, Nombre: "Pieza 2", Descripcion: "Habitación con vista al jardín"},
	{ID: 3, Nombre: "Pieza 3", Descripcion: "Habitación estándar"},
}

var productos = []Producto{
	{ID: 1, Nombre: "Champaña", Precio: 30.5},
	{ID: 2, Nombre: "Cerveza", Precio: 15.0},
	{ID: 3, Nombre: "Snacks", Precio: 8.0},
}

// Handlers de la API
func getPiezas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(piezas)
}

func getProductos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(productos)
}

func main() {
	http.HandleFunc("/api/piezas", getPiezas)
	http.HandleFunc("/api/productos", getProductos)

	http.ListenAndServe(":8085", nil)
}

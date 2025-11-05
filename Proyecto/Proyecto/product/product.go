package models

type Producto struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

type Movimiento struct {
	ID         int    `json:"id"`
	ProductoID int    `json:"producto_id"`
	Tipo       string `json:"tipo"`
	Cantidad   int    `json:"cantidad"`
	Razon      string `json:"razon"`
}

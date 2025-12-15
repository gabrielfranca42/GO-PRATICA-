package model

type Product struct{
	ID     int `json:"id_procuct"`
	Name   string `json:"name"`
	Price  float64 `json:"price"`
}
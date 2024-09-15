package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func productToJSON(p Product) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func jsonToProduct(jsonStr string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func main() {
	product := Product{Name: "Laptop", Price: 999.99, Quantity: 10}

	jsonStr, err := productToJSON(product)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	fmt.Println("Product in JSON format:", jsonStr)

	decodedProduct, err := jsonToProduct(jsonStr)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}

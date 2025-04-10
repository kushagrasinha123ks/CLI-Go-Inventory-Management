package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Name     string
	Category string
	Quantity int
	Price    float64
}

var inventory = make(map[string]Product)

const dataFile = "inventory.json"

func saveInventory() {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error saving inventory:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(inventory); err != nil {
		fmt.Println("Error encoding inventory:", err)
	}
}

func loadInventory() {
	file, err := os.Open(dataFile)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&inventory); err != nil {
		fmt.Println("Error decoding inventory:", err)
	}
}

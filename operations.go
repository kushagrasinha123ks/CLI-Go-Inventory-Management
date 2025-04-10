package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func addProduct(reader *bufio.Reader) {
	fmt.Print("Enter product name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter product category: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Enter quantity: ")
	quantityStr, _ := reader.ReadString('\n')
	quantity, _ := strconv.Atoi(strings.TrimSpace(quantityStr))

	fmt.Print("Enter price: ")
	priceStr, _ := reader.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

	inventory[name] = Product{Name: name, Category: category, Quantity: quantity, Price: price}
	fmt.Println("Product added successfully.")
}

func viewInventory(reader *bufio.Reader) {
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Println("View Inventory:")
	fmt.Println("1. Sort by Category")
	fmt.Println("2. Sort by Price")
	fmt.Print("Enter your choice: ")
	sortChoice, _ := reader.ReadString('\n')
	sortChoice = strings.TrimSpace(sortChoice)

	products := make([]Product, 0, len(inventory))
	for _, product := range inventory {
		products = append(products, product)
	}

	if sortChoice == "1" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	} else if sortChoice == "2" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
	}

	for _, product := range products {
		fmt.Printf("Name: %s, Category: %s, Quantity: %d, Price: %.2f\n", product.Name, product.Category, product.Quantity, product.Price)
	}
}

func updateProduct(reader *bufio.Reader) {
	fmt.Print("Enter product name to update: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	product, exists := inventory[name]
	if !exists {
		fmt.Println("Product not found.")
		return
	}

	fmt.Print("Enter new quantity (or press enter to skip): ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = strings.TrimSpace(quantityStr)
	if quantityStr != "" {
		quantity, _ := strconv.Atoi(quantityStr)
		product.Quantity = quantity
	}

	fmt.Print("Enter new price (or press enter to skip): ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	if priceStr != "" {
		price, _ := strconv.ParseFloat(priceStr, 64)
		product.Price = price
	}

	inventory[name] = product
	fmt.Println("Product updated successfully.")
}

func deleteProduct(reader *bufio.Reader) {
	fmt.Print("Enter product name to delete: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	_, exists := inventory[name]
	if !exists {
		fmt.Println("Product not found.")
		return
	}

	delete(inventory, name)
	fmt.Println("Product deleted successfully.")
}

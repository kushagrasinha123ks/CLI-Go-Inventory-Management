package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateCSVReport() {
	now := time.Now()
	filename := fmt.Sprintf("inventory_report_%s.csv", now.Format("20060102150405"))

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Name", "Category", "Quantity", "Price"}
	writer.Write(header)

	for _, product := range inventory {
		row := []string{product.Name, product.Category, strconv.Itoa(product.Quantity), fmt.Sprintf("%.2f", product.Price)}
		writer.Write(row)
	}

	fmt.Printf("CSV report generated: %s\n", filename)
}

func importCSV(reader *bufio.Reader) {
	fmt.Print("Enter CSV file path: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	if len(records) > 0 {
		for _, record := range records[1:] {
			if len(record) == 4 {
				name := record[0]
				category := record[1]
				quantity, err := strconv.Atoi(record[2])
				if err != nil {
					fmt.Println("Invalid quantity in CSV:", err)
					continue
				}
				price, err := strconv.ParseFloat(record[3], 64)
				if err != nil {
					fmt.Println("Invalid price in CSV:", err)
					continue
				}

				newProduct := Product{Name: name, Category: category, Quantity: quantity, Price: price}
				existingProduct, exists := inventory[name]

				if exists {
					if existingProduct != newProduct {
						fmt.Printf("Product %s already exists. Choose an option:\n", name)
						fmt.Println("1. Keep both (add with a suffix)")
						fmt.Println("2. Update with new data")
						fmt.Print("Enter your choice: ")
						choice, _ := reader.ReadString('\n')
						choice = strings.TrimSpace(choice)

						if choice == "1" {
							suffix := 1
							for {
								newName := fmt.Sprintf("%s (%d)", name, suffix)
								if _, exists := inventory[newName]; !exists {
									inventory[newName] = newProduct
									fmt.Printf("Product %s added as %s.\n", name, newName)
									break
								}
								suffix++
							}

						} else if choice == "2" {
							inventory[name] = newProduct
							fmt.Printf("Product %s updated from CSV.\n", name)
						} else {
							fmt.Println("Invalid choice. Skipping.")
						}

					} else {
						fmt.Printf("Product %s already exists and is identical. Skipping.\n", name)
					}
				} else {
					inventory[name] = newProduct
					fmt.Printf("Product %s added from CSV.\n", name)
				}
			} else {
				fmt.Println("Invalid CSV row:", record)
			}
		}
		fmt.Println("CSV data import completed.")
	} else {
		fmt.Println("CSV file is empty.")
	}
}

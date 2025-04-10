package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	loadInventory()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. View Inventory")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Generate CSV Report")
		fmt.Println("6. Import CSV")
		fmt.Println("7. Exit")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addProduct(reader)
		case "2":
			viewInventory(reader)
		case "3":
			updateProduct(reader)
		case "4":
			deleteProduct(reader)
		case "5":
			generateCSVReport()
		case "6":
			importCSV(reader)
		case "7":
			saveInventory()
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

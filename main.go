package main

import (
	"fmt"

	product "github.com/gilwong00/go-invoice-cli/pkg"
)

const NAME string = "Invoice Generator"

func main() {
	fmt.Print(NAME + "\n\n")
	products := []product.Product{}
	for {
		prod := product.CreateProduct()
		products = append(products, prod)
		fmt.Println("Product added. Press q to exit or any other key to continue")
		var nextInput string
		fmt.Scanln(&nextInput)
		if nextInput == "q" {
			break
		}
	}
	var totalAmount float64
	var invoiceContent string
	for i, prod := range products {
		invoiceContent += fmt.Sprintf("%v)\t Product: \"%v\" \t Price: \"%v\" \t Quantity: %v \n", i+1, prod.Name, prod.Price, prod.Quantity)
		totalAmount += prod.Price * prod.Quantity
	}
	invoiceContent += fmt.Sprintf("\nTotal cost: %v\n", totalAmount)
	fmt.Println("\n", invoiceContent)
	var shouldGenerate string
	fmt.Println("Would you like to generate an invoice?")
	fmt.Println("Press y to generate or any key to cancel")
	fmt.Scanln(&shouldGenerate)
	if shouldGenerate == "y" {
		err := product.GenerateInvoice(NAME, invoiceContent)
		if err != nil {
			fmt.Println("Failed to generate invoice", err)
		}
		fmt.Println("Invoice generated!")
	}
}

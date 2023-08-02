package product

import (
	"fmt"
	"os"
	"time"
)

type Product struct {
	Name     string
	Price    float64
	Quantity float64
}

func CreateProduct() Product {
	product := Product{}

	// get product from user
	fmt.Println("Please enter the name of the product")
	fmt.Scan(&product.Name)

	// get product price
	fmt.Println("How much does this cost?")
	fmt.Scan(&product.Price)

	// get quantity
	fmt.Println("How many were purchased?")
	fmt.Scan(&product.Quantity)
	return product
}

func GenerateInvoice(appName string, data string) error {
	fileName := fmt.Sprintf("invoice-%v.txt", time.Now())
	content := fmt.Sprintf("%v\n", appName)
	content += fmt.Sprintf("Date: %v\n\n", time.Now())
	content += data
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

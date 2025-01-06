package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

type Inventory struct {
	Products []Product
}

func (inv *Inventory) AddProduct(id, name, priceStr string, stock int) error {
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return errors.New("invalid price format")
	}
	inv.Products = append(inv.Products, Product{ID: id, Name: name, Price: price, Stock: stock})
	return nil
}

func (inv *Inventory) UpdateStock(id string, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range inv.Products {
		if product.ID == id {
			inv.Products[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

func (inv *Inventory) SearchProduct(query string) (*Product, error) {
	for _, product := range inv.Products {
		if strings.EqualFold(product.Name, query) || product.ID == query {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (inv *Inventory) DisplayInventory(sortBy string) {
	if sortBy == "price" {
		sort.Slice(inv.Products, func(i, j int) bool {
			return inv.Products[i].Price < inv.Products[j].Price
		})
	} else if sortBy == "stock" {
		sort.Slice(inv.Products, func(i, j int) bool {
			return inv.Products[i].Stock < inv.Products[j].Stock
		})
	}

	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inv.Products {
		fmt.Printf("%-10s %-20s rs %-9.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func main() {
	inventory := Inventory{}

	inventory.AddProduct("101", "Laptop", "47000.99", 10)
	inventory.AddProduct("102", "Smartphone", "35000.49", 25)
	inventory.AddProduct("103", "Air Purifier", "15000.99", 15)
	inventory.AddProduct("104", "AC", "32000.99", 14)

	err := inventory.UpdateStock("104", 20)
	if err != nil {
		fmt.Println("Error updating stock:", err)
	}

	product, err := inventory.SearchProduct("Laptop")
	if err != nil {
		fmt.Println("Error searching product:", err)
	} else {
		fmt.Printf("Found Product: %+v\n", *product)
	}

	fmt.Println("\nInventory sorted by price:")
	inventory.DisplayInventory("price")

	fmt.Println("\nInventory sorted by stock:")
	inventory.DisplayInventory("stock")
}

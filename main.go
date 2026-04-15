package main

import "fmt"

type Customer struct {
	ID   int
	Name string
	Age  int
}

func main() {

	customers := []Customer{
		{1, "Sreeja", 25},
		{2, "Bharath", 30},
		{3, "Sai", 28},
	}

	for _, c := range customers {
		fmt.Printf("ID: %d , Name: %s, Age : %d\n", c.ID, c.Name, c.Age)
	}
}

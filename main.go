package main

import "fmt"

type customer struct {
	ID   int
	Name string
	Age  int
}

func main() {
	customers := []customer{
		{1, "Sreeja", 25},
		{2, "Bharath", 16},
		{3, "Sai", 28},
		{4, "Raj", 10},
		{5, "Ravi", 18},
	}
	/*for _, c := range customers {
		fmt.Printf("ID: %d , Name: %s , Age: %d\n", c.ID, c.Name, c.Age)
	}*/
	fmt.Printf("%-5s %-10s %-5s %-40s\n", "ID", "Name", "Age", "Status")
	for _, c := range customers {
		if c.Age < 12 {
			fmt.Printf("%-5d %-10s %-5d %-40s\n", c.ID, c.Name, c.Age, "Not Allowed")
		} else if c.Age > 12 && c.Age < 18 {
			fmt.Printf("%-5d %-10s %-5d %-40s\n", c.ID, c.Name, c.Age, "Allowed but must come with someone above 21")
		} else {
			fmt.Printf("%-5d %-10s %-5d %-40s\n", c.ID, c.Name, c.Age, "Allowed")
		}
	}

}

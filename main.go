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
		{2, "Bharath", 12},
		{3, "Sai", 28},
		{4, "Raj", 10},
		{5, "Ravi", 18},
	}
	ActualTicketPrice := 15.0
	fmt.Printf("%-5s %-10s %-5s %-45s %-10s\n", "ID", "Name", "Age", "Status", "TicketPrice")
	for _, c := range customers {
		if c.Age < 18 {
			TicketPrice := ActualTicketPrice / 2
			if c.Age < 12 {
				fmt.Printf("%-5d %-10s %-5d %-45s %-10s\n", c.ID, c.Name, c.Age, "Not Allowed", "NA")
			} else if c.Age >= 12 && c.Age < 18 {
				fmt.Printf("%-5d %-10s %-5d %-45s %-10.2f\n", c.ID, c.Name, c.Age, "Allowed but must come with someone above 21", TicketPrice)
			}
		} else {
			fmt.Printf("%-5d %-10s %-5d %-45s %-10.2f\n", c.ID, c.Name, c.Age, "Allowed", ActualTicketPrice)

		}
	}

}

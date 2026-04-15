package main

import "fmt"

type customer struct {
	ID        int
	Name      string
	Age       int
	IsLoyalty bool
}

func main() {
	customers := []customer{
		{1, "Sreeja", 25, true},
		{2, "Bharath", 12, false},
		{3, "Sai", 28, true},
		{4, "Raj", 10, false},
		{5, "Ravi", 18, true},
	}
	actualTicketPrice := 15.0
	fmt.Printf("%-5s %-10s %-5s %-45s %-10s\n", "ID", "Name", "Age", "Status", "TicketPrice")
	for _, c := range customers {
		status, TicketPrice := calculateTicketPrice(c, actualTicketPrice)
		if TicketPrice == 0 {
			fmt.Printf("%-5d %-10s %-5d %-45s %-10s\n", c.ID, c.Name, c.Age, status, "NA")
		} else {
			fmt.Printf("%-5d %-10s %-5d %-45s %-10.2f\n", c.ID, c.Name, c.Age, status, TicketPrice)
		}
	}

}

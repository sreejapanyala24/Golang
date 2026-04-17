package main

import (
	"fmt"
	"net/http"
)

type customer struct {
	ID        int
	Name      string
	Age       int
	IsLoyalty bool
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		customers := []customer{
			{1, "Sreeja", 25, true},
			{2, "Bharath", 12, false},
			{3, "Sai", 28, true},
			{4, "Raj", 10, false},
			{5, "Ravi", 15, true},
			{6, "Sarayu", 4, false},
		}

		actualTicketPrice := 15.0

		fmt.Fprintf(w, "%-5s %-10s %-5s %-10s %-45s %-20s %-42s %-15s\n",
			"ID", "Name", "Age", "IsLoyalty", "Status", "ActualTicketPrice", "OfferApplied", "FinalTicketPrice")

		for _, c := range customers {
			status, finalTicketPrice, offer := calculateTicketPrice(c, actualTicketPrice)

			if finalTicketPrice == 0 {
				fmt.Fprintf(w, "%-5d %-10s %-5d %-10v %-45s $%-19.2f %-42s %-15s\n",
					c.ID, c.Name, c.Age, c.IsLoyalty, status, actualTicketPrice, offer, "NA")
			} else {
				fmt.Fprintf(w, "%-5d %-10s %-5d %-10v %-45s $%-19.2f %-42s $%-14.2f\n",
					c.ID, c.Name, c.Age, c.IsLoyalty, status, actualTicketPrice, offer, finalTicketPrice)
			}
		}

	})
	fmt.Println("Server running on port 9090...")
	http.ListenAndServe(":9090", nil)
}

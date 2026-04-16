package main

func calculateTicketPrice(c customer, actualTicketPrice float64) (string, float64, string) {

	if c.Age < 12 {
		return "Not Allowed", 0, "No offer"
	}

	TicketPrice := actualTicketPrice
	offers := []string{}

	if c.Age >= 12 && c.Age < 18 {
		TicketPrice = actualTicketPrice / 2
		offers = append(offers, "50% Teen Discount")
	}

	if c.IsLoyalty {
		TicketPrice = TicketPrice - (TicketPrice * 0.10)
		offers = append(offers, "10% Loyalty Discount")
	}

	offerApplied := "No offer"
	if len(offers) > 0 {
		offerApplied = ""
		for i, o := range offers {
			if i > 0 {
				offerApplied += " + "
			}
			offerApplied += o
		}
	}

	if c.Age >= 12 && c.Age < 18 {
		return "Allowed but must come with someone above 18", TicketPrice, offerApplied
	}

	return "Allowed", TicketPrice, offerApplied
}

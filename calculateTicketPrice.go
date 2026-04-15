package main

func calculateTicketPrice(c customer, actualTicketPrice float64) (string, float64) {
	if c.Age < 12 {
		return "Not Allowed", 0
	}
	TicketPrice := actualTicketPrice

	if c.Age >= 12 && c.Age < 18 {
		TicketPrice = actualTicketPrice / 2
	}
	if c.IsLoyalty {
		TicketPrice = TicketPrice - (TicketPrice * 0.10)
	}
	if c.Age >= 12 && c.Age < 18 {
		return "Allowed but must come with someone above 21", TicketPrice
	}
	return "Allowed", TicketPrice
}

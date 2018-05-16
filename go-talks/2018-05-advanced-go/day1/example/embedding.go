package main

// START 1 OMIT
type Discount struct {
	rate float64
}
type LoyaltyDiscount struct {
	Discount
	additionalRate float64
}

// END 1 OMIT

func (d Discount) Calculate(amount float64) float64 {
	return amount - (amount*d.rate)/100.0
}

// START 3 OMIT
func (d LoyaltyDiscount) Calculate(amount float64) float64 {
	baseDiscount := d.Discount.Calculate(amount)
	return baseDiscount - (baseDiscount*d.additionalRate)/100.0
}

// END 3 OMIT

func main() {

	// START 2 OMIT
	LoyaltyDiscount{
		Discount: Discount{
			rate: 50.0,
		},
		additionalRate: 10.0,
	}
	// END 2 OMIT

}

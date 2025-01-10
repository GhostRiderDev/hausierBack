package models

// Home defines a Home property to sell
type Home struct {
	title       string
	description string
	address     Address
	price       float64
}

// Address define a Location of a home
type Address struct {
	Location string
	Country  string
	State    string
	City     string
}

// Title returns a title reference from a home
func (h *Home) Title() *string {
	return &h.title
}

// Description returns a description reference from a home
func (h *Home) Description() *string {
	return &h.description
}

// Address returns a address location reference from a home
func (h *Home) Address() *Address {
	return &h.address
}
// Address returns a price reference from a home
func (h *Home) Price() *float64 {
	return &h.price
}

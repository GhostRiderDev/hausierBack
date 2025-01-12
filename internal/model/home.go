package model

// Home defines a Home property to sell
type Home struct {
	Title       string
	Description string
	Address     Address
	Price       float64
}

// Address define a Location of a home
type Address struct {
	Location string
	Country  string
	State    string
	City     string
}
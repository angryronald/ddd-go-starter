package model

type CustomerApplicationModel struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Addresses []*AddressApplicationModel
}

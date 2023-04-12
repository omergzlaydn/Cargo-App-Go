package structure

type Customer struct {
	IdNumber    int
	Name        string
	LastName    string
	PhoneNumber int
	Address     string
}

func NewCustomer(IdNumber int, Name, LastName string, PhoneNumber int, Address string) Customer {
	return Customer{
		IdNumber:    IdNumber,
		Name:        Name,
		LastName:    LastName,
		PhoneNumber: PhoneNumber,
		Address:     Address,
	}
}

//Customer Adress Update
func (customer *Customer) SetCustomerAddress(address string) {
	customer.Address = address
}

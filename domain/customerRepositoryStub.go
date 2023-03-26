package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "ABC1", City: "City1", Zipcode: "ZIPCODE1", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "ABC2", City: "City2", Zipcode: "ZIPCODE2", DateofBirth: "2000-01-02", Status: "2"},
		{Id: "1003", Name: "ABC3", City: "City3", Zipcode: "ZIPCODE3", DateofBirth: "2000-01-03", Status: "3"},
	}
	return CustomerRepositoryStub{customers: customers}
}

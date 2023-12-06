package domain

import "fmt"

type Customer struct {
	ID    string
	Name  string
	Email string
}

func (c Customer) Print() {
	fmt.Printf("Customer[ID: %s, Name: %s, Email: %s]\n", c.ID, c.Name, c.Email)
}

type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetByID(string) (Customer, error)
	GetAll() ([]Customer, error)
}

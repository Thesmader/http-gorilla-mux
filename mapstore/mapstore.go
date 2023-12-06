package mapstore

import (
	"errors"

	"assignment3/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{
		store: make(map[string]domain.Customer, 0),
	}
}

func (m *MapStore) Create(c domain.Customer) error {
	if _, ok := m.store[c.ID]; ok {
		return errors.New("customer already exists")
	}
	m.store[c.ID] = c
	return nil
}

func (m *MapStore) Update(id string, c domain.Customer) error {
	if _, ok := m.store[id]; !ok {
		return errors.New("customer does not exist")
	}
	m.store[id] = c
	return nil
}

func (m *MapStore) Delete(id string) error {
	if _, ok := m.store[id]; !ok {
		return errors.New("customer does not exist")
	}
	delete(m.store, id)
	return nil
}

func (m *MapStore) GetByID(id string) (domain.Customer, error) {
	if _, ok := m.store[id]; !ok {
		return domain.Customer{}, errors.New("customer does not exist")
	}
	return m.store[id], nil
}

func (m *MapStore) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	for _, customer := range m.store {
		customers = append(customers, customer)
	}
	return customers, nil
}

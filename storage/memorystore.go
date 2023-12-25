package storage

import (
	"rest-book/model"
	"fmt"
	"strconv"
)

type MemoryStore struct {
	Restaurants []*model.Restaurant
}

func NewMemoryStore() Store{
	return &MemoryStore{
	}
}
func (m *MemoryStore) GetRestaurantDetails(id string) (*model.Restaurant, error) {
	idx, err := strconv.Atoi(id)
	if idx > len(m.Restaurants) - 1 {
		return nil, fmt.Errorf("no restaurant")
	}
	if err != nil {
		return nil, err
	}
	return  m.Restaurants[idx], nil
}

func(m *MemoryStore) AddRestaurantDetails(rest *model.Restaurant) error {
	m.Restaurants = append(m.Restaurants, rest)
	return nil
}

func(m *MemoryStore) UpdateRestaurantDetails(id string, rest *model.Restaurant) error {
	return nil
}

func(m *MemoryStore) GetAllRestaurantDetails() ([]*model.Restaurant, error) {
	return m.Restaurants, nil
}

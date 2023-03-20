package datas

import (
	"pkgs/internal/repositories/interfaces"
)

type InMemoryRepository struct {
	mapData map[string]string

	dataInterface interfaces.Repository
}

func New() *InMemoryRepository {
	return &InMemoryRepository{mapData: make(map[string]string)}
}

func (s *InMemoryRepository) Create(key string, value string) bool {
	s.mapData[key] = value

	return true
}

func (s *InMemoryRepository) Read(key string) string {
	return s.mapData[key]
}

func (s *InMemoryRepository) Update(key string, value string) bool {
	s.mapData[key] = value

	return true
}

func (s *InMemoryRepository) Delete(key string) bool {
	delete(s.mapData, key)

	return true
}

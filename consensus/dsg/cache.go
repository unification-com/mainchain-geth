package dsg

type Cache struct {
	Validations map[uint64]map[uint64]bool `json:"validations"`
}

func NewCache() *Cache {
	cache := &Cache{
		Validations: map[uint64]map[uint64]bool{},
	}
	return cache
}

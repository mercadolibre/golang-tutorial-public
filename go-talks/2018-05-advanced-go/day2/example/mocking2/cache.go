package mocking2

// Service ...
type Service struct {
	// memcached *memcache.Client // Dificil de testear
	cache Cache
}

var _ Cache = &Service{}

// New ...
func New(c Cache) *Service {
	return &Service{
		cache: c,
	}
}

// CachePut ...
type CachePut interface {
	Put(string, []byte) error
}

// CacheGet ...
type CacheGet interface {
	Get(string) ([]byte, error)
}

// Cache ...
type Cache interface {
	CacheGet
	CachePut
}

// Put ...
func (s *Service) Put(key string, value []byte) error {
	return s.cache.Put(key, value)
}

// Get ...
func (s *Service) Get(key string) ([]byte, error) {
	return s.cache.Get(key)
}

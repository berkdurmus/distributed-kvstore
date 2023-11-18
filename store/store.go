package store

type Store struct {
	data map[string]string
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Put(key, value string) {
	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

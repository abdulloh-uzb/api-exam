package repo

type InMemorystorageI interface {
	SetWithTTL(key, value string, seconds int) error
	Get(key string) (interface{}, error)
	Exists(key string) (interface{}, error)
}

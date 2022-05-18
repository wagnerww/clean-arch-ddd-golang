package shared

type RepositoryCacheInterface[T any] interface {
	Create(keyId string, entity T) error
	Find(keyId string) (T, error)
	Delete(keyId string) error
}

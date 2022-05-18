package shared

type RepositoryInterface[T any] interface {
	Create(entity T) error
	Update(entity T) error
	FindById(id string) (T, error)
	FindAll() ([]T, error)
}

package shared

type RepositoryInterface[T any] interface {
	Create(entity T) error
	Update(entity T) error
	FindById(id string) (entity T, err error)
	FindAll() (entities []T, err error)
}

package repository

type Repository[T any] interface {
	Get(id uint) T
	Create(t T) (uint, error)
	All() []T
	Update(t T) (T, error)
	Delete(id uint) error
}

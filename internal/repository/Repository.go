package repository

// M = Domain Model
// R = Repo Model (Gorm model)
type Repository[M any, R any] interface {
	Get(id uint) (M, error)
	Create(t R) (M, error)
	All() []M
	Update(t R) (M, error)
	Delete(id uint) error
}

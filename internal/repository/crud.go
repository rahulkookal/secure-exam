package repository

type CRUDE[T any] interface {
	Create(item *T) error
	GetByID(id string) (*T, error)
	Update(id string, item *T) error
	Delete(id string) error
}

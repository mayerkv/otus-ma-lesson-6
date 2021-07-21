package domain

type UserRepository interface {
	Save(user *User) error
	FindById(id int) (*User, error)
	NextId() (int, error)
	Delete(user *User) error
}

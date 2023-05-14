package repository

type User struct {
	ID       int
	Username string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
}

type userRepository struct {
	users []*User
}

// Create implements UserRepository
func (repo *userRepository) Create(user *User) error {
	append(repo.users, user)
	return nil
}

// Delete implements UserRepository
func (repo *userRepository) Delete(user *User) error {
	panic("unimplemented")
}

// GetAll implements UserRepository
func (repo *userRepository) GetAll() ([]*User, error) {
	return repo.users, nil
}

// GetByID implements UserRepository
func (repo *userRepository) GetByID(id int) (*User, error) {
	return repo.users[id], nil
}

// Update implements UserRepository
func (*userRepository) Update(user *User) error {
	panic("unimplemented")
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: []*User{},
	}
}

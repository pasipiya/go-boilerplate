package repository

import "github.com/pasipiya/go-boilerplate/internal/model"

type UserRepository struct {
	users []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []model.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		},
	}
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	return r.users, nil
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *UserRepository) Create(name, email string) (*model.User, error) {
	newUser := model.NewUser(len(r.users)+1, name, email)
	r.users = append(r.users, *newUser)
	return newUser, nil
}

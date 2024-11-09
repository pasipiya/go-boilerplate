package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(id int, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

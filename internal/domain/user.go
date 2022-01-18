package domain

import "context"

type UserRepository interface {
	Find(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Save(context.Context, *User) error
}

type User struct {
	ID            int      `json:"id"`
	Email         string   `json:"email"`
	AuthTokenHash string   `json:"-"`
	PasswordHash  string   `json:"-"`
	Name          string   `json:"name"`
	Surname       string   `json:"surname"`
	Age           int      `json:"age"`
	Sex           string   `json:"sex"`
	Interests     []string `json:"interests"`
	City          string   `json:"city"`
	Friends       []int    `json:"friends"`
}

func NewUser(
	email string,
	passwordHash string,
	name string,
	surname string,
	age int,
	sex string,
	interests []string,
	city string,
) *User {
	return &User{
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		Surname:      surname,
		Age:          age,
		Sex:          sex,
		Interests:    interests,
		City:         city,
	}
}

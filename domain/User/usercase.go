package user

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"ecommer.com/ecommerce/model"
	"github.com/google/uuid"
)

type User struct {
	storage Storage
}

// new user
func new(s Storage) User {
	return User{storage: s}
}

// create user
func (u User) Create(m *model.User) error {

	ID, err := uuid.NewUUID()

	if err != nil {
		return fmt.Errorf("error while creating uuid: %w", err)
	}

	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("error while generating password: %w", err)
	}

	m.Password = string(password)

	if m.Deatils == nil {
		m.Deatils = []byte("{}")
	}

	m.CreatedAt = time.Now().Unix()

	u.storage.Create(m)

	if err != nil {
		return fmt.Errorf("error while creating user: %w", err)
	}

	m.Password = ""

	return nil
}

// get user by email

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("error while getting user by email: %w", err)
	}

	return user, nil
}

// get all users
func (u User) GetAll() (model.Users, error) {
	// get all users
	users, err := u.storage.GetAll()
	//
	if err != nil {
		return model.Users{}, fmt.Errorf("error while getting all users: %w", err)
	}

	return users, nil
}

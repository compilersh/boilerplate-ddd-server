package repository

import "fmt"

type User struct {
	ID       string
	Username string
}

type UserRepository interface {
	CreateUser(user User) (*User, error)
	GetUser(id string) (User, error)
	GetAllUsers() ([]User, error)
}

// Make sure inmemdb implements UserRepository
var _ UserRepository = &InMemDB{}

func (db *InMemDB) CreateUser(user User) (*User, error) {
	for _, u := range db.users {
		if u.Username == user.Username {
			return nil, fmt.Errorf("username %s already exists", user.Username)
		}
	}
	db.users = append(db.users, user)
	return &user, nil
}

func (db *InMemDB) GetUser(id string) (User, error) {
	for _, u := range db.users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with id %s not found", id)
}

func (db *InMemDB) GetAllUsers() ([]User, error) {
	return db.users, nil
}

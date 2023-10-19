package user

import (
	"fmt"
	"math/rand"

	"github.com/compilersh/boilerplate-ddd-server/repository"
)

type Config struct {
	SomeValue string
}

// User represents a user in the system.
// We map this to the User struct in the repository package.
// This can be useful if we want to change the User struct in the repository without
// having to change the User struct in the service.
type User struct {
	ID       string
	Username string
}

func FromRepository(user repository.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
	}
}

func (u User) ToRepository() repository.User {
	return repository.User{
		ID:       u.ID,
		Username: u.Username,
	}
}

func (u User) ToUserRes() userRes {
	return userRes{
		ID:       u.ID,
		Username: u.Username,
	}
}

type UserService struct {
	someValue string
	users     repository.UserRepository
}

func NewUserService(cfg Config, users repository.UserRepository) *UserService {
	return &UserService{
		someValue: cfg.SomeValue,
		users:     users,
	}
}

func (s *UserService) CreateUser(userReq UserReq) (User, error) {
	user := &User{
		ID:       someRandomID(),
		Username: userReq.Username,
	}
	repoUser, err := s.users.CreateUser(user.ToRepository())
	if err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}
	return FromRepository(*repoUser), nil
}

func (s *UserService) GetUser(id string) (User, error) {
	user, err := s.users.GetUser(id)
	if err != nil {
		return User{}, fmt.Errorf("get user: %w", err)
	}

	return FromRepository(user), nil
}

func (s *UserService) GetAllUsers() ([]User, error) {
	users, err := s.users.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}

	var result []User
	for _, user := range users {
		result = append(result, FromRepository(user))
	}

	return result, nil
}

// this is just a random string generator
// don't use this in production
func someRandomID() string {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = s[rand.Intn(len(s))]
	}
	return string(b)
}

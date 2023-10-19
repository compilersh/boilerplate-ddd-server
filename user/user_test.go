package user_test

import (
	"testing"

	"github.com/compilersh/boilerplate-ddd-server/repository"
	"github.com/compilersh/boilerplate-ddd-server/user"
)

func TestCreateUser(t *testing.T) {
	t.Run("valid user", func(t *testing.T) {
		userName := "testuser"

		testDB := &TestUserDB{}

		service := user.NewUserService(user.Config{
			SomeValue: "testval",
		}, testDB)

		user, err := service.CreateUser(user.UserReq{
			Username: userName,
		})
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if user.Username != userName {
			t.Errorf("expected username %s, got %s", userName, user.Username)
		}

	})
}

// create a mock db that implements the user interface
type TestUserDB struct{}

// make sure TestUserDB implements the UserRepository interface
var _ repository.UserRepository = &TestUserDB{}

func (db *TestUserDB) CreateUser(user repository.User) (*repository.User, error) {
	return &user, nil
}

func (db *TestUserDB) GetUser(id string) (repository.User, error) {
	return repository.User{}, nil
}

func (db *TestUserDB) GetAllUsers() ([]repository.User, error) {
	return nil, nil
}

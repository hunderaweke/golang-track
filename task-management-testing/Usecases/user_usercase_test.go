package usecases

import (
	"testing"
	domain "testing-api/Domain"
	infrastructure "testing-api/Infrastructure"
	"testing-api/mocks"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	userRepository mocks.UserRepository
	userUsecase    domain.UserUsecase
	users          []domain.User
}

func (suite *UserUsecaseTestSuite) SetupSuite() {
	suite.userRepository = *mocks.NewUserRepository(suite.T())
	suite.userUsecase = NewUserUsecase(&suite.userRepository, time.Duration(4*time.Second), context.TODO())
	suite.users = []domain.User{
		{
			ID:       "1",
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Password: "password123",
		},
		{
			ID:       "2",
			Name:     "Jane Smith",
			Email:    "janesmith@example.com",
			Password: "securepass",
		},
		{
			ID:       "3",
			Name:     "Alice Johnson",
			Email:    "alicej@example.com",
			Password: "mypassword",
		},
		{
			ID:       "4",
			Name:     "Bob Brown",
			Email:    "bobb@example.com",
			Password: "adminpass",
		},
	}
}

func (suite *UserUsecaseTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userRepository.On("Create", mock.Anything, u).Return(&u, nil)
	}
	for _, u := range suite.users {
		createdUser, err := suite.userUsecase.Create(u)
		assert.NoError(err)
		assert.Equal(u, *createdUser)
	}
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestGet() {
	assert := assert.New(suite.T())
	suite.userRepository.On("Get", mock.Anything).Return(suite.users, nil)
	users, err := suite.userUsecase.Get()
	assert.NoError(err)
	assert.Equal(suite.users, users)
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestGetByID() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userRepository.On("GetByID", mock.Anything, u.ID).Return(&u, nil)
	}
	for _, u := range suite.users {
		createdUser, err := suite.userUsecase.GetByID(u.ID)
		assert.NoError(err)
		assert.Equal(u, *createdUser)
	}
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestGetByEmail() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userRepository.On("GetByEmail", mock.Anything, u.Email).Return(&u, nil)
	}
	for _, u := range suite.users {
		createdUser, err := suite.userUsecase.GetByEmail(u.Email)
		assert.NoError(err)
		assert.Equal(u, *createdUser)
	}
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUpdate() {
	assert := assert.New(suite.T())

	for _, u := range suite.users {
		suite.userRepository.On("Update", mock.Anything, u.ID, u).Return(&u, nil)
	}
	for _, u := range suite.users {
		updatedUser, err := suite.userUsecase.Update(u.ID, u)
		assert.NoError(err)
		assert.Equal(u, *updatedUser)
	}
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestDelete() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userRepository.On("Delete", mock.Anything, u.ID).Return(nil)
	}
	for _, u := range suite.users {
		err := suite.userUsecase.Delete(u.ID)
		assert.NoError(err)
	}
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestLogin() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		hashedPassword, _ := infrastructure.HashPassword(u.Password)
		exsitingUser := domain.User{ID: u.ID, Password: hashedPassword, Email: u.Email, IsAdmin: u.IsAdmin}
		suite.userRepository.On("GetByEmail", mock.Anything, u.Email).Return(&exsitingUser, nil)
	}
	for _, u := range suite.users {
		foundUser, err := suite.userUsecase.Login(u)
		assert.NoError(err)
		assert.Equal(u, foundUser)
	}
	suite.userRepository.AssertExpectations(suite.T())
}

func TestUserUsecaseSuite(t *testing.T) {
	tSuite := new(UserUsecaseTestSuite)
	suite.Run(t, tSuite)
}

package repository

import (
	"context"
	"testing"

	domain "testing-api/Domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mongoMocks "github.com/sv-tools/mongoifc/mocks/mockery"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	repository domain.UserRepository
	mockCol    *mongoMocks.Collection
	mockDb     *mongoMocks.Database
	data       []domain.User
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	suite.mockCol = new(mongoMocks.Collection)
	suite.mockDb = new(mongoMocks.Database)

	suite.mockCol.On("InsertOne", context.Background(), mock.Anything).Return(
		&mongo.InsertOneResult{
			InsertedID: primitive.NewObjectID(),
		},
		nil,
	).Once()

	suite.mockCol.On("Find", context.Background(), mock.Anything).Return(
		new(mongoMocks.Cursor),
		nil,
	).Once()

	suite.mockCol.On("UpdateOne", context.Background(), mock.Anything, mock.Anything).Return(
		&mongo.UpdateResult{
			MatchedCount:  1,
			ModifiedCount: 1,
		},
		nil,
	).Once()

	suite.mockCol.On("DeleteOne", context.Background(), mock.Anything).Return(
		&mongo.DeleteResult{
			DeletedCount: 1,
		},
		nil,
	).Once()

	suite.mockDb.On("Collection", domain.UserCollection).Return(suite.mockCol)

	suite.repository = NewUserService(context.TODO(), suite.mockDb)

	suite.data = []domain.User{
		{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Password: "password123",
		},
		{
			Name:     "Jane Smith",
			Email:    "janesmith@example.com",
			Password: "securepass",
		},
		{
			Name:     "Alice Johnson",
			Email:    "alicej@example.com",
			Password: "mypassword",
		},
		{
			Name:     "Bob Brown",
			Email:    "bobb@example.com",
			Password: "adminpass",
		},
	}
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	suite.mockCol.AssertExpectations(suite.T())
	suite.mockDb.AssertExpectations(suite.T())
}

func (suite *UserRepositoryTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	for i, u := range suite.data {
		created, err := suite.repository.Create(context.TODO(), u)
		assert.NoError(err)
		if i == 0 {
			assert.True(created.IsAdmin, "expected for the first user to be an admin")
		}
		assert.Equal(created.Name, u.Name)
		assert.Equal(created.Email, u.Email)
		assert.Equal(created.Password, u.Password)
		suite.data[i] = *created
	}
}

func (suite *UserRepositoryTestSuite) TestGet() {
	assert := assert.New(suite.T())
	users, err := suite.repository.Get(context.TODO())
	assert.NoError(err)
	assert.Equal(len(users), len(suite.data))
}

func (suite *UserRepositoryTestSuite) TestGetByID() {
	assert := assert.New(suite.T())
	for _, u := range suite.data {
		got, err := suite.repository.GetByID(context.TODO(), u.ID)
		assert.NoError(err)
		assert.Equal(u.Name, got.Name)
		assert.Equal(u.ID, got.ID)
		assert.Equal(u.IsAdmin, got.IsAdmin)
		assert.Equal(u.Email, got.Email)
	}
}

func (suite *UserRepositoryTestSuite) TestGetByEmail() {
	assert := assert.New(suite.T())
	for _, u := range suite.data {
		got, err := suite.repository.GetByEmail(context.TODO(), u.Email)
		assert.NoError(err)
		assert.Equal(u.Name, got.Name)
		assert.Equal(u.ID, got.ID)
		assert.Equal(u.IsAdmin, got.IsAdmin)
		assert.Equal(u.Email, got.Email)
	}
}

func (suite *UserRepositoryTestSuite) TestUserDelete() {
	assert := assert.New(suite.T())
	id := suite.data[1].ID
	err := suite.repository.Delete(context.TODO(), id)
	assert.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestPromote() {
	assert := assert.New(suite.T())
	id := suite.data[2].ID
	err := suite.repository.PromoteUser(context.TODO(), id)
	assert.NoError(err)
	user, err := suite.repository.GetByID(context.TODO(), id)
	assert.True(user.IsAdmin)
}

func (suite *UserRepositoryTestSuite) TestUpdate() {
	assert := assert.New(suite.T())
	for i, u := range suite.data {
		data := domain.User{Password: u.Password + "dfjskd", Email: u.Email + "sdjfks"}
		updated, err := suite.repository.Update(context.TODO(), u.ID, data)
		assert.NoError(err)
		assert.Equal(data.Password, updated.Password)
		assert.Equal(data.Email, updated.Email)
		suite.data[i] = *updated
	}
}

func TestUserRepositorySuite(t *testing.T) {
	tSuite := new(UserRepositoryTestSuite)
	suite.Run(t, tSuite)
}

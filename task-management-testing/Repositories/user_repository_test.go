package repository

import (
	"context"
	"log"
	"os"
	"testing"
	domain "testing-api/Domain"
	"testing-api/database"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	repository domain.UserRepository
	db         *mongo.Database
	data       []domain.User
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	dbUri := os.Getenv("MONGODB_URL")
	clnt, err := database.NewConnection(context.TODO(), dbUri)
	if err != nil {
		log.Fatal(err)
	}
	db := clnt.Database("task_management_api_test")
	suite.repository = NewUserService(context.TODO(), db)
	suite.db = db
	db.Collection(domain.UserCollection).DeleteMany(context.TODO(), bson.D{{}}, options.Delete())
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
	collection := suite.db.Collection(domain.UserCollection)
	collection.DeleteMany(context.TODO(), bson.D{{}}, options.Delete())
	suite.db.Drop(context.TODO())
	suite.db.Client().Disconnect(context.TODO())
}

func (suite *UserRepositoryTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	for i, t := range suite.data {
		created, err := suite.repository.Create(context.TODO(), t)
		assert.NoError(err)
		if i == 0 {
			assert.True(created.IsAdmin, "expected for the first user to be and admin")
		}
		assert.Equal(created.Name, t.Name)
		assert.Equal(created.Email, t.Email)
		assert.Equal(created.Password, t.Password)
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

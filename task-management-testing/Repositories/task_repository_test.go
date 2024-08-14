package repository

import (
	"context"
	"testing"
	domain "testing-api/Domain"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mongoMocks "github.com/sv-tools/mongoifc/mocks/mockery"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	tasks          []domain.Task
	taskRepository domain.TaskRepository
	mockCol        *mongoMocks.Collection
	mockDb         *mongoMocks.Database
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
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

	suite.mockDb.On("Collection", domain.TaskCollection).Return(suite.mockCol)
	suite.taskRepository = NewTaskService(context.TODO(), suite.mockDb)
	suite.tasks = []domain.Task{
		{
			UserID:      "user123",
			Title:       "Complete project report",
			Description: "Finish the final report and submit it to the manager.",
			DueDate:     time.Date(2024, 8, 15, 12, 0, 0, 0, time.UTC),
			Status:      "pending",
		},
		{
			UserID:      "user123",
			Title:       "Prepare presentation",
			Description: "Create slides for the upcoming conference.",
			DueDate:     time.Date(2024, 8, 20, 9, 0, 0, 0, time.UTC),
			Status:      "pending",
		},
		{
			UserID:      "user456",
			Title:       "Update software documentation",
			Description: "Revise the user manual and update the API docs.",
			DueDate:     time.Date(2024, 8, 18, 17, 0, 0, 0, time.UTC),
			Status:      "done",
		},
		{
			UserID:      "user789",
			Title:       "Plan team meeting",
			Description: "Organize and schedule the next team meeting.",
			DueDate:     time.Date(2024, 8, 22, 10, 0, 0, 0, time.UTC),
			Status:      "pending",
		},
		{
			UserID:      "user101",
			Title:       "Review code submissions",
			Description: "Go through the latest code reviews and provide feedback.",
			DueDate:     time.Date(2024, 8, 19, 15, 0, 0, 0, time.UTC),
			Status:      "done",
		},
	}
}

func (suite *TaskRepositoryTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	for i, t := range suite.tasks {
		createdTask, err := suite.taskRepository.Create(context.TODO(), t)
		assert.NoError(err)
		assert.Equal(t.UserID, createdTask.UserID)
		assert.Equal(t.DueDate, createdTask.DueDate)
		assert.Equal(t.Status, createdTask.Status)
		assert.Equal(t.Title, createdTask.Title)
		assert.Equal(t.Description, createdTask.Description)
		suite.tasks[i].ID = createdTask.ID
	}
}

func (suite *TaskRepositoryTestSuite) TestGet() {
	assert := assert.New(suite.T())
	got, err := suite.taskRepository.Get(context.TODO())
	assert.NoError(err)
	for i, t := range suite.tasks {
		assert.Equal(t.UserID, got[i].UserID)
		assert.Equal(t.DueDate, got[i].DueDate)
		assert.Equal(t.Status, got[i].Status)
		assert.Equal(t.Title, got[i].Title)
		assert.Equal(t.Description, got[i].Description)
		suite.tasks[i].ID = got[i].ID
	}
}

func (suite *TaskRepositoryTestSuite) TestGetByID() {
	assert := assert.New(suite.T())
	for _, t := range suite.tasks {
		got, err := suite.taskRepository.GetByID(context.TODO(), t.ID)
		assert.NoError(err)
		assert.Equal(got, t)
	}
}

func (suite *TaskRepositoryTestSuite) TestGetByUserID() {
	taskGroup := map[string][]domain.Task{}
	assert := assert.New(suite.T())
	for _, t := range suite.tasks {
		taskGroup[t.UserID] = append(taskGroup[t.UserID], t)
	}
	for userID, expectedTasks := range taskGroup {
		got, err := suite.taskRepository.GetByUserID(context.TODO(), userID)
		assert.NoError(err)
		assert.Equal(got, expectedTasks)
	}
}

func (suite *TaskRepositoryTestSuite) TestUpdate() {
	assert := assert.New(suite.T())
	for i, t := range suite.tasks {
		updateData := t
		updateData.Description += "Edited"
		got, err := suite.taskRepository.Update(context.TODO(), t.ID, updateData)
		assert.NoError(err)
		assert.Equal(updateData, *got)
		suite.tasks[i] = updateData
	}
}

func (suite *TaskRepositoryTestSuite) TestDelete() {
	assert := assert.New(suite.T())
	err := suite.taskRepository.Delete(context.TODO(), suite.tasks[0].ID)
	assert.NoError(err)
	suite.tasks = suite.tasks[1:]
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	suite.mockCol.AssertExpectations(suite.T())
	suite.mockDb.AssertExpectations(suite.T())
}

func TestTaskRepositorySuite(t *testing.T) {
	tSuite := new(TaskRepositoryTestSuite)
	suite.Run(t, tSuite)
}

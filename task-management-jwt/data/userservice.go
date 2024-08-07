package data

import (
	"context"
	"errors"
	"fmt"
	"task-management-api-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	collection *mongo.Collection
	Count      int
}

type UserTasks []models.Task

func NewUserService(db *mongo.Database) *UserService {
	c := db.Collection("users")
	emailIndexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	c.Indexes().CreateOne(context.TODO(), emailIndexModel)
	cnt, _ := c.CountDocuments(context.Background(), bson.D{{}}, options.Count())
	return &UserService{collection: c, Count: int(cnt)}
}

func (u *UserService) GetUsers() ([]models.User, error) {
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	users := []models.User{}
	c, err := u.collection.Find(context.Background(), bson.D{{}}, opts)
	if err != nil {
		return users, err
	}
	for c.Next(context.TODO()) {
		var e models.User
		err := c.Decode(&e)
		if err != nil {
			return users, err
		}
		users = append(users, e)
	}
	return users, nil
}

func (u *UserService) GetUserByID(userID string) (*models.User, error) {
	opts := options.FindOne()
	user := models.User{}
	res := u.collection.FindOne(context.TODO(), bson.D{{"id", userID}}, opts)
	err := res.Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserService) GetByEmail(email string) (*models.User, error) {
	opts := options.FindOne()
	user := models.User{}
	res := u.collection.FindOne(context.TODO(), bson.D{{"email", email}}, opts)
	err := res.Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserService) UpdateUser(userID string, data models.User) (*models.User, error) {
	user, err := u.GetUserByID(userID)
	if err != nil {
		return user, err
	}
	if data.Name != "" {
		user.Name = data.Name
	}
	if data.Email != "" {
		user.Email = data.Email
	}
	opts := options.Replace()
	res, err := u.collection.ReplaceOne(context.Background(), bson.D{{"id", userID}}, user, opts)
	if err != nil {
		return user, err
	}
	if res.ModifiedCount != 1 {
		return user, fmt.Errorf("error modifing users more than one update")
	}
	return user, nil
}

func (u *UserService) DeleteUser(userID string) error {
	opts := options.Delete()
	res, err := u.collection.DeleteOne(context.Background(), bson.D{{"id", userID}}, opts)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("user with id %v not found", userID)
	}
	if res.DeletedCount != 1 {
		return fmt.Errorf("deleted %v users", res.DeletedCount)
	}
	return nil
}

func (u *UserService) Create(user *models.User) (*models.User, error) {
	cnt, _ := u.collection.CountDocuments(context.Background(), bson.D{{}}, options.Count())
	if cnt == 0 {
		user.IsAdmin = true
	} else if cnt > 0 && user.IsAdmin {
		return user, errors.New("promoting user requires admin access")
	}
	_, err := u.collection.InsertOne(context.Background(), user, options.InsertOne())
	return user, err
}

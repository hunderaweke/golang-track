package repository

import (
	domain "clean-architecture/Domain"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	collection *mongo.Collection
}

type UserTasks []domain.Task

func NewUserService(c context.Context, db *mongo.Database) domain.UserRepository {
	collection := db.Collection(domain.UserCollection)
	emailIndexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	collection.Indexes().CreateOne(c, emailIndexModel)
	return &UserService{collection: collection}
}

func (u *UserService) Get(c context.Context) ([]domain.User, error) {
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	users := []domain.User{}
	cursor, err := u.collection.Find(c, bson.D{{}}, opts)
	if err != nil {
		return users, err
	}
	for cursor.Next(c) {
		var e domain.User
		err := cursor.Decode(&e)
		if err != nil {
			return users, err
		}
		users = append(users, e)
	}
	return users, nil
}

func (u *UserService) GetByID(c context.Context, userID string) (*domain.User, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return &domain.User{}, fmt.Errorf("user with id %v not found", userID)
	}
	opts := options.FindOne()
	user := domain.User{}
	res := u.collection.FindOne(c, bson.D{{"_id", id}}, opts)
	err = res.Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserService) GetByEmail(c context.Context, email string) (*domain.User, error) {
	opts := options.FindOne()
	user := domain.User{}
	res := u.collection.FindOne(c, bson.D{{"email", email}}, opts)
	err := res.Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserService) PromoteUser(c context.Context, userID string) error {
	user, err := u.GetByID(c, userID)
	if err != nil {
		return err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	user.IsAdmin = true
	opts := options.Replace()
	res, err := u.collection.ReplaceOne(c, bson.D{{"id", id}}, user, opts)
	if err != nil {
		return err
	}
	if res.ModifiedCount != 1 {
		return fmt.Errorf("error modifing: users more than one update")
	}
	return nil
}

func (u *UserService) Update(c context.Context, userID string, data domain.User) (*domain.User, error) {
	user, err := u.GetByID(c, userID)
	if err != nil {
		return user, err
	}
	if data.Name != "" {
		user.Name = data.Name
	}
	if data.Email != "" {
		user.Email = data.Email
	}
	if data.IsAdmin != false {
		user.IsAdmin = data.IsAdmin
	}
	opts := options.Replace()
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return &domain.User{}, err
	}
	res, err := u.collection.ReplaceOne(c, bson.D{{"_id", id}}, user, opts)
	if err != nil {
		return user, err
	}
	if res.ModifiedCount != 1 {
		return user, fmt.Errorf("error modifing users more than one update")
	}
	return user, nil
}

func (u *UserService) Delete(c context.Context, userID string) error {
	opts := options.Delete()
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("user with id %v not found", userID)
	}
	res, err := u.collection.DeleteOne(c, bson.D{{"_id", id}}, opts)
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

func (u *UserService) Create(c context.Context, user domain.User) (*domain.User, error) {
	cnt, _ := u.collection.CountDocuments(c, bson.D{{}}, options.Count())
	if cnt == 0 {
		user.IsAdmin = true
	} else if cnt > 0 && user.IsAdmin {
		return &user, errors.New("promoting user requires admin access")
	}
	user.ID = primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	_, err := u.collection.InsertOne(context.Background(), user, options.InsertOne())
	if mongo.IsDuplicateKeyError(err) {
		return &user, fmt.Errorf("user with the same email already exists")
	}
	return &user, err
}

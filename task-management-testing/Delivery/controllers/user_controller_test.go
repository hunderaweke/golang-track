package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	domain "testing-api/Domain"
	infrastructure "testing-api/Infrastructure"
	"testing-api/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	router         *gin.Engine
	userUsecase    mocks.UserUsecase
	userController UserController
	users          []domain.User
	admin          domain.User
}

func (suite *UserControllerTestSuite) SetupSuite() {
	suite.router = gin.Default()
	os.Setenv("JWT_SECRET", "testsecret")
	suite.admin = domain.User{
		ID:      "1",
		IsAdmin: true,
		Email:   "admin@admin.com",
	}
	suite.userUsecase = *mocks.NewUserUsecase(suite.T())
	suite.userController = *NewUserController(&suite.userUsecase)
}

func (suite *UserControllerTestSuite) TestCreateController() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userUsecase.On("Create", u).Return(u, nil)
	}
	suite.router.POST("/register", infrastructure.JWTMiddleware(), infrastructure.AdminMiddleware(), suite.userController.Create)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, u := range suite.users {
		data, _ := json.Marshal(u)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(data))
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusCreated, w.Code)
		assert.JSONEq(string(data), w.Body.String())
	}
}

func (suite *UserControllerTestSuite) TestGetUsers() {
	assert := assert.New(suite.T())
	suite.userUsecase.On("Get").Return(suite.users, nil)
	suite.router.GET("/users", infrastructure.JWTMiddleware(), suite.userController.GetUsers)
	token, _ := infrastructure.GenerateToken(suite.admin)
	data, _ := json.Marshal(suite.users)
	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)
	assert.JSONEq(string(data), w.Body.String())
}

func (suite *UserControllerTestSuite) TestGetUserByID() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userUsecase.On("GetByID", u.ID).Return(u, nil)
	}
	suite.router.GET("/users/:id", infrastructure.JWTMiddleware(), suite.userController.GetUserByID)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, u := range suite.users {
		data, _ := json.Marshal(u)
		req, _ := http.NewRequest("GET", "/users/"+u.ID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusOK, w.Code)
		assert.JSONEq(string(data), w.Body.String())
	}
}

func (suite *UserControllerTestSuite) TestUpdateUser() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userUsecase.On("Update", u.ID, u).Return(u, nil)
		suite.userUsecase.On("GetByID", u.ID).Return(u, nil)
	}
	suite.router.PUT("/users/:id", infrastructure.JWTMiddleware(), infrastructure.AdminMiddleware(), suite.userController.UpdateUser)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, u := range suite.users {
		data, _ := json.Marshal(u)
		req, _ := http.NewRequest("PUT", "/users/"+u.ID, bytes.NewBuffer(data))
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusOK, w.Code)
		assert.JSONEq(string(data), w.Body.String())
	}
}

func (suite *UserControllerTestSuite) TestDeleteTask() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userUsecase.On("Delete", u.ID).Return(nil)
		suite.userUsecase.On("GetByID", u.ID).Return(u, nil)
	}
	suite.router.DELETE("/users/:id", infrastructure.JWTMiddleware(), infrastructure.AdminMiddleware(), suite.userController.DeleteUser)
	token, _ := infrastructure.GenerateToken(suite.admin)
	for _, u := range suite.users {
		req, _ := http.NewRequest("DELETE", "/users/"+u.ID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusNoContent, w.Code)
	}
}

func (suite *UserControllerTestSuite) TestLogin() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userUsecase.On("Login", u).Return(u, nil)
	}
	suite.router.POST("/login/", suite.userController.Login)
	for _, u := range suite.users {
		data, err := json.Marshal(u)
		assert.NoError(err)
		req, err := http.NewRequest("POST", "/login/", bytes.NewBuffer(data))
		assert.NoError(err)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(http.StatusOK, w.Code)
		expectedToken, err := infrastructure.GenerateToken(u)
		assert.NoError(err)
		expectedResponse := struct {
			Message string      `json:"message"`
			Token   string      `json:"token "`
			User    domain.User `json:"user"`
		}{
			Message: "successful login",
			Token:   expectedToken,
			User:    u,
		}
		expectedResponseJson, err := json.Marshal(expectedResponse)
		assert.NoError(err)
		assert.JSONEq(string(expectedResponseJson), w.Body.String())
	}
}

func (suite *UserControllerTestSuite) TestPromoteUser() {
	assert := assert.New(suite.T())
	for _, u := range suite.users {
		suite.userUsecase.On("Promote", u.ID).Return(nil)
	}
	suite.router.POST("/promote/", infrastructure.JWTMiddleware(), infrastructure.AdminMiddleware(), suite.userController.PromoteUser)
	for _, u := range suite.users {
		data, err := json.Marshal(u)
		assert.NoError(err)
		req, err := http.NewRequest("POST", "/promote/", bytes.NewBuffer(data))
		assert.NoError(err)
		token, err := infrastructure.GenerateToken(suite.admin)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		expectedResponse := struct {
			Message string `json:"message"`
		}{Message: "user promoted"}
		expectedResponseJson, err := json.Marshal(expectedResponse)
		assert.Equal(http.StatusAccepted, w.Code)
		assert.JSONEq(string(expectedResponseJson), w.Body.String())
	}
}

func TestUserController(t *testing.T) {
	tSuite := new(UserControllerTestSuite)
	suite.Run(t, tSuite)
}

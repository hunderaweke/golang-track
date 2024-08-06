package controllers

import (
	"net/http"
	"task-management-api-mongodb/data"
	"task-management-api-mongodb/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	userService *data.UserService
	nextID      int
}

func NewUserController(db *mongo.Database) *UserController {
	u := data.NewUserService(db)
	return &UserController{userService: u, nextID: u.Count + 1}
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
}

func (u *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := u.userService.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedUser)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := u.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (u *UserController) GetUserTasks(c *gin.Context) {
}

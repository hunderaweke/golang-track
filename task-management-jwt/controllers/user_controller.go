package controllers

import (
	"net/http"
	"os"
	"strconv"
	"task-management-api-mongodb/data"
	"task-management-api-mongodb/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserController) Create(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = strconv.Itoa(u.nextID)
	u.nextID++
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	_, err = u.userService.Create(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (u *UserController) Login(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	exsitingUser, err := u.userService.GetByEmail(user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(exsitingUser.Password), []byte(user.Password)) != nil {
		c.JSON(500, gin.H{"error": "invalid password or email"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": exsitingUser.ID,
		"email":   exsitingUser.Email,
		"name":    exsitingUser.Name,
	})
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful login", "token": jwtToken})
}

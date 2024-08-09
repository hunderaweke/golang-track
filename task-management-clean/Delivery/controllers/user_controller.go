package controllers

import (
	domain "clean-architecture/Domain"
	infrastructure "clean-architecture/Infrastructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.userUsecase.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (u *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (u *UserController) PromoteUser(c *gin.Context) {
	var user *domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if user.ID == "" && user.Email == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid repository"})
		return
	}
	if user.Email != "" && user.ID == "" {
		user, err = u.userUsecase.GetByEmail(user.Email)
	}
	user.IsAdmin = true
	err = u.userUsecase.PromoteUser(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user promoted to admin"})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := domain.User{}
	userClaims := getUserClaims(c)
	if userClaims.UserID != id || !userClaims.IsAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "method not allowed"})
		return
	}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.IsAdmin {
		c.JSON(http.StatusNotModified, gin.H{"error": "promoting user requires admin access"})
		return
	}
	updatedUser, err := u.userUsecase.Update(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedUser)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userClaims := getUserClaims(c)
	if !userClaims.IsAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deleting user requires admin access"})
		return
	}
	if err := u.userUsecase.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func validateUser(u domain.User) error {
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

func (u *UserController) Create(c *gin.Context) {
	var user struct {
		Password string `json:"password"`
		domain.User
	}
	err := c.BindJSON(&user)
	user.User.Password = user.Password
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validateUser(user.User)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	hashPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
	}
	user.User.Password = hashPassword
	newUser, err := u.userUsecase.Create(user.User)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newUser)
}

func (u *UserController) Login(c *gin.Context) {
	var user struct {
		Password string `json:"password"`
		domain.User
	}
	err := c.Bind(&user)
	user.User.Password = user.Password
	if err != nil {
		c.JSON(500, gin.H{"error": "user entity is required"})
		return
	}
	err = validateUser(user.User)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
	}
	exsitingUser, err := u.userUsecase.Login(user.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtToken, err := infrastructure.GenerateToken(exsitingUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful login", "token": jwtToken, "user": exsitingUser})
}

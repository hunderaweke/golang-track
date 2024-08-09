package controllers

import (
	domain "clean-architecture/Domain"
	infrastructure "clean-architecture/Infrastructure"
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
	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
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

func (u *UserController) Create(c *gin.Context) {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
	}
	user.Password = hashPassword
	_, err = u.userUsecase.Create(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (u *UserController) Login(c *gin.Context) {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "user entity is required"})
		return
	}
	exsitingUser, err := u.userUsecase.GetByEmail(user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if !infrastructure.ComparePassword(user.Password, exsitingUser.Password) {
		c.JSON(500, gin.H{"error": "invalid password or email"})
		return
	}
	jwtToken, err := infrastructure.GenerateToken(domain.User(*exsitingUser))
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful login", "token": jwtToken, "user": exsitingUser})
}

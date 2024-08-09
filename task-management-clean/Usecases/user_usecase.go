package usecases

import (
	domain "clean-architecture/Domain"
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository  domain.UserRepository
	timeoutDuration time.Duration
	ctx             context.Context
}

func NewUserUsecase(userRepository domain.UserRepository, timeoutDuration time.Duration, ctx context.Context) domain.UserUsecase {
	return &userUsecase{
		userRepository:  userRepository,
		timeoutDuration: timeoutDuration,
		ctx:             ctx,
	}
}

func (uu *userUsecase) Create(u domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(uu.ctx, uu.timeoutDuration)
	defer cancel()
	return uu.userRepository.Create(ctx, u)
}

func (uu *userUsecase) Get() ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(uu.ctx, uu.timeoutDuration)
	defer cancel()
	return uu.userRepository.Get(ctx)
}

func (uu *userUsecase) GetByID(userID string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(uu.ctx, uu.timeoutDuration)
	defer cancel()
	return uu.userRepository.GetByID(ctx, userID)
}

func (uu *userUsecase) GetByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(uu.ctx, uu.timeoutDuration)
	defer cancel()
	return uu.userRepository.GetByEmail(ctx, email)
}

func (uu *userUsecase) PromoteUser(userID string) error {
	ctx, cancel := context.WithTimeout(uu.ctx, uu.timeoutDuration)
	defer cancel()
	return uu.userRepository.PromoteUser(ctx, userID)
}

func (uu *userUsecase) Login(user domain.User) (domain.User, error) {
	exsitingUser, err := uu.GetByEmail(user.Email)
	if err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(exsitingUser.Password), []byte(user.Password)) != nil {
		return user, fmt.Errorf("invalid email or password")
	}
	return *exsitingUser, nil
}

func (uu *userUsecase) Delete(userID string) error {
	ctx, cancel := context.WithTimeout(uu.ctx, uu.timeoutDuration)
	defer cancel()
	return uu.userRepository.Delete(ctx, userID)
}

func (uu *userUsecase) Update(userID string, data domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(uu.ctx, uu.timeoutDuration)
	defer cancel()
	return uu.userRepository.Update(ctx, userID, data)
}
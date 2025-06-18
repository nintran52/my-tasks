package usecase

import (
	"errors"

	"github.com/nintran52/my-tasks/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	Repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) *UserUsecase {
	return &UserUsecase{Repo: r}
}

func (uc *UserUsecase) CreateUser(u *domain.User) error {
	// Check if email exists
	existing, err := uc.Repo.GetByEmail(u.Email)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("email already exists")
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	// Save to database
	u.Password = string(hashedPassword)
	return uc.Repo.Create(u)
}

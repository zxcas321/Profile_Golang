package service

import (
	"errors"

	"github.com/google/uuid"
	"profile_go/internal/model"
	"profile_go/internal/repository"
	"profile_go/internal/request"
	"profile_go/internal/response"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAll() ([]response.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []response.UserResponse
	for _, u := range users {
		result = append(result, response.NewUserResponse(u.ID, u.Email))
	}
	return result, nil
}

func (s *UserService) GetByID(id string) (*response.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	res := response.NewUserResponse(user.ID, user.Email)
	return &res, nil
}

func (s *UserService) Create(req request.CreateUserRequest) error {
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil {
		return errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &model.User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		PasswordHash: string(hashed),
	}
	return s.userRepo.Create(user)
}

func (s *UserService) Update(id string, req request.UpdateUserRequest) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	if req.Email != "" && req.Email != user.Email {
		existing, _ := s.userRepo.FindByEmail(req.Email)
		if existing != nil {
			return errors.New("email already used")
		}
		user.Email = req.Email
	}

	if req.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash password")
		}
		user.PasswordHash = string(hashed)
	}

	return s.userRepo.Update(user)
}


func (s *UserService) Delete(id string) error {
	return s.userRepo.SoftDelete(id)
}
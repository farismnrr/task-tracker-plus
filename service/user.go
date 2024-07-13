/**
 * Package service provides interfaces and implementations for user management.
 * 
 * Interfaces:
 * 
 * - UserService: Interface defining methods for user management.
 *   Methods:
 *   - Register: Method to register a new user.
 *   - Login: Method to authenticate and generate JWT token for a user.
 *   - GetUserTaskCategory: Method to retrieve user task categories.
 * 
 * Structs:
 * 
 * - userService: Struct implementing the UserService interface.
 *   Fields:
 *   - userRepo: Instance of repo.UserRepository for user repository operations.
 *   - sessionsRepo: Instance of repo.SessionRepository for session repository operations.
 *   Methods:
 *   - NewUserService: Function to create a new instance of userService.
 *   - Register: Method to register a new user by checking email existence, creating a new user, and storing user session.
 *   - Login: Method to authenticate a user by email and password, generate JWT token, and manage user session.
 *   - GetUserTaskCategory: Method to retrieve user task categories using the user repository.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo     repo.UserRepository
	sessionsRepo repo.SessionRepository
}

func NewUserService(userRepository repo.UserRepository, sessionsRepo repo.SessionRepository) UserService {
	return &userService{userRepository, sessionsRepo}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if dbUser.Email == "" || dbUser.ID == 0 {
		return nil, errors.New("user not found")
	}

	if user.Password != dbUser.Password {
		return nil, errors.New("wrong email or password")
	}

	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &model.Claims{
		Email: dbUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	session := model.Session{
		Token:  tokenString,
		Email:  user.Email,
		Expiry: expirationTime,
	}

	if _, err := s.sessionsRepo.SessionAvailEmail(session.Email); err != nil {
		s.sessionsRepo.AddSessions(session)
	} else {
		s.sessionsRepo.UpdateSessions(session)
	}

	return &tokenString, nil
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	return s.userRepo.GetUserTaskCategory()
}

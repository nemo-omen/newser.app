package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"newser.app/internal/domain/entity"
	"newser.app/internal/dto"
	"newser.app/internal/infra/repository"
	"newser.app/shared"
)

var (
	defaultUserCollections = []string{"read", "unread", "saved"}
)

type AuthService struct {
	authRepo repository.AuthRepository
	// collectionRepo repository.CollectionRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return AuthService{
		authRepo: authRepo,
	}
}

func (s *AuthService) Login(email, password string) (*entity.User, error) {
	storedHash, err := s.authRepo.GetHashedPasswordByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Invalid password",
			"Login",
			"value.Password",
		)
	}

	u, err := s.authRepo.FindByEmail(email)
	if err != nil {
		// TODO: rewrite so AppError originates in repo
		return nil, shared.NewAppError(
			err,
			fmt.Sprintf("%s is not registered", email),
			"Login",
			"value.Email",
		)
	}
	return u, nil
}

// Register creates a new user in the database.
// It also creates the default collections for the user.
func (s *AuthService) Register(email, name, hashedPassword string) (*dto.UserDTO, error) {
	// Initializing a new User entity will
	// validate the email and name through
	// value objects.
	user, err := entity.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	// Create a UserDAO from the User entity
	// and the hashed password.
	ud, err := dto.NewUserDAO(user.ID.String(), user.Name.String(), user.Email.String(), hashedPassword)
	if err != nil {
		return nil, err
	}

	// Create the default collections for the user.
	collections := make([]*entity.Collection, 0)
	for _, title := range defaultUserCollections {
		c, err := entity.NewCollection(title, user.ID.String())
		if err != nil {
			return nil, err
		}
		collections = append(collections, c)
	}

	// Create the user and the default collections
	// in a single transaction.
	err = s.authRepo.CreateUser(*ud, collections)
	if err != nil {
		return nil, err
	}

	// Create the user DTO to send back to handler.
	udto := dto.UserDTO{}.FromDomain(user)

	return &udto, nil
}

func (s *AuthService) Logout() error {
	return nil
}

func (s *AuthService) GetUserById(userId int64) (*entity.User, error) {
	return nil, nil
}

func (s *AuthService) GetUserByEmail(email string) (*entity.User, error) {
	return nil, nil
}

package service

import (
	"backend/api/repository"
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

// UserService UserService struct
type UserService struct {
	repository repository.UserRepository
}

// NewUserService : returns the UserService struct instance
func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

// hashPassword -> returns hashed password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash -> func to check hashed pwd
// compares password & hashed pwd
func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Register -> hashes password and calls user repo save method
func (u UserService) Register(user models.User) error {
	hashedPwd, err := hashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPwd
	return u.repository.Save(user)
}

// Save -> calls user repository save method
func (u UserService) Save(user models.User) error {
	return u.repository.Save(user)
}

// FindAll -> calls to repo FindAll method
func (u UserService) FindAll(user models.User, keyword string) (*[]models.User, int64, error) {
	return u.repository.FindAll(user, keyword)
}

// Update -> calls userRepo update method
func (u UserService) Update(user models.User) error {
	return u.repository.Update(user)
}

// Delete -> calls user repo delete method
func (u UserService) Delete(id int64) error {
	var user models.User
	user.ID = id
	return u.repository.Delete(user)
}

// Find -> calls user repo find method
func (u UserService) Find(user models.User) (models.User, error) {
	return u.repository.Find(user)
}

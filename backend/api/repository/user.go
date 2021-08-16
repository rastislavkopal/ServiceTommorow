package repository

import (
	"backend/common"
	"backend/models"
)

// UserRepository -> UserRepository
type UserRepository struct {
	db common.Database
}

// new Repository : fetching database
func NewUserRepository(db common.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// Save -> method for saving user into DB
func (u UserRepository) Save(user models.User) error {
	return u.db.DB.Create(&user).Error
}

// FindAll -> Method for fetching all users from db
func (u UserRepository) FindAll(user models.User, keyword string) (*[]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0

	queryBuilder := u.db.DB.Order("created_at desc").Model(&models.User{})

	// search param -> first_name
	if keyword != "" {
		queryWord := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			u.db.DB.Where("user.first_name LIKE ? ", queryWord))

	}
	err := queryBuilder.Where(user).Find(&users).Count(&totalRows).Error

	return &users, totalRows, err
}

// Update -> method for update
func (u UserRepository) Update(user models.User) error {
	return u.db.DB.Save(&user).Error
}

// Find -> method for fetching user by ID
func (u UserRepository) Find(user models.User) (models.User, error) {
	var users models.User
	err := u.db.DB.Model(&models.User{}).Where(&user).Take(&users).Error
	return users, err
}

// Find -> method for fetching user by email
func (u UserRepository) FindByEmail(user models.User) (models.User, error) {
	var usr models.User
	err := u.db.DB.Where("email = ?", user.Email).First(&usr).Error
	return usr, err
}

// Delete -> method to delete user by id
func (u UserRepository) Delete(user models.User) error {
	return u.db.DB.Delete(&user).Error
}

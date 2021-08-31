package repository

import (
	"backend/common"
	"backend/models"
)

// TaskRepository -> TaskRepository
type TaskRepository struct {
	db *common.Database
}

// new Repository : fetching database
func NewTaskRepository(db *common.Database) TaskRepository {
	return TaskRepository{
		db: db,
	}
}

// Save -> method for saving task into DB
func (t *TaskRepository) Save(task models.Task) error {
	return t.db.DB.Create(&task).Error
}

// FindAll -> Method for fetching all users from db
func (t *TaskRepository) FindAll(task models.Task, keyword string) (*[]models.Task, int64, error) {
	var tasks []models.Task
	var totalRows int64 = 0

	queryBuilder := t.db.DB.Order("created_at desc").Model(&models.Task{})

	// search param -> title
	if keyword != "" {
		queryWord := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			t.db.DB.Where("task.Title LIKE ? ", queryWord))

	}
	err := queryBuilder.Where(task).Find(&tasks).Count(&totalRows).Error

	return &tasks, totalRows, err
}

// Update -> method for update
func (t *TaskRepository) Update(task models.Task) error {
	return t.db.DB.Save(&task).Error
}

// Find -> method for fetching task by ID
func (t *TaskRepository) Find(task models.Task) (models.Task, error) {
	var tasks models.Task
	err := t.db.DB.Model(&models.Task{}).Where(&task).Take(&tasks).Error
	return tasks, err
}

// Delete -> method to delete task by id
func (u *TaskRepository) Delete(task models.Task) error {
	return u.db.DB.Delete(&task).Error
}

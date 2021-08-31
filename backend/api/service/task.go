package service

import (
	"backend/api/repository"
	"backend/models"
)

// TaskService TaskService struct
type TaskService struct {
	repository *repository.TaskRepository
}

// NewTaskService : returns the TaskService struct instance
func NewTaskService(r *repository.TaskRepository) TaskService {
	return TaskService{
		repository: r,
	}
}

// Save -> calls task repository save method
func (t *TaskService) Save(task models.Task) error {
	return t.repository.Save(task)
}

// FindAll -> calls to repo FindAll method
func (t *TaskService) FindAll(task models.Task, keyword string) (*[]models.Task, int64, error) {
	return t.repository.FindAll(task, keyword)
}

// Update -> calls TaskRepo update method
func (t *TaskService) Update(task models.Task) error {
	return t.repository.Update(task)
}

// Delete -> calls task repo delete method
func (t *TaskService) Delete(id uint64) error {
	var task models.Task
	task.ID = id
	return t.repository.Delete(task)
}

// Find -> calls task repo find method
func (t *TaskService) Find(task models.Task) (models.Task, error) {
	return t.repository.Find(task)
}

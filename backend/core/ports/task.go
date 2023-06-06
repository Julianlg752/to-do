package ports

import "todo/core/models"

type TaskPortInterface interface {
	GetSingleTask(taskID int64) (*models.Task, error)
	GetAllTask() ([]*models.Task, error)
	UpdateTask(taskRequest *models.Task) error
	DeleteTask(taskID int64) error
	AddTask(taskRequest *models.Task) (*models.Task, error)
}

type TaskRepositoryInterface interface {
	GetSingleTask(taskID int64) (*models.Task, error)
	GetAllTask() ([]*models.Task, error)
	UpdateTask(taskRequest *models.Task) error
	DeleteTask(taskID int64) error
	AddTask(taskRequest *models.Task) (*models.Task, error)
}

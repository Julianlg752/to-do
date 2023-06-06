package interactor

import (
	"todo/core/models"
	"todo/core/ports"
	"todo/core/repository"

	"github.com/ansel1/merry/v2"
)

type Task struct {
	taskRepository ports.TaskRepositoryInterface
}

var TaskInteractor ports.TaskPortInterface = &Task{
	taskRepository: repository.TaskRepository,
}

func (t *Task) GetSingleTask(taskID int64) (*models.Task, error) {
	return t.taskRepository.GetSingleTask(taskID)
}

func (t *Task) GetAllTask() ([]*models.Task, error) {
	return t.taskRepository.GetAllTask()
}

func (t *Task) UpdateTask(taskRequest *models.Task) error {
	if err := taskRequest.Validate(); err != nil {
		return merry.Wrap(err)
	}
	return merry.Wrap(t.taskRepository.UpdateTask(taskRequest))
}

func (t *Task) DeleteTask(taskID int64) error {
	return merry.Wrap(t.taskRepository.DeleteTask(taskID))
}

func (t *Task) AddTask(taskRequest *models.Task) (*models.Task, error) {
	if err := taskRequest.Validate(); err != nil {
		return nil, merry.Wrap(err)
	}
	var err error
	taskRequest, err = t.taskRepository.AddTask(taskRequest)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	return taskRequest, nil
}

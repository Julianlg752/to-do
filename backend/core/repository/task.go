package repository

import (
	"time"
	"todo/core/models"
	"todo/core/ports"
	"todo/datastore"

	"github.com/ansel1/merry/v2"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var taskTable = goqu.T("tasks")

type taskRepository struct {
	db        *goqu.Database
	taskTable exp.IdentifierExpression
}

var TaskRepository ports.TaskRepositoryInterface = &taskRepository{
	db:        &datastore.GoquDB,
	taskTable: taskTable,
}

func (t *taskRepository) GetSingleTask(taskID int64) (*models.Task, error) {
	var task models.Task
	ok, err := t.db.From(t.taskTable).Where(
		t.taskTable.Col("id").Eq(taskID),
	).ScanStruct(&task)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	if !ok {
		return nil, merry.New("record not found")
	}

	return &task, nil
}

func (t *taskRepository) GetAllTask() ([]*models.Task, error) {
	var taskList []*models.Task
	//TODO: add offset to pagination
	if err := t.db.From(t.taskTable).ScanStructs(&taskList); err != nil {
		return nil, merry.Wrap(err)
	}

	return taskList, nil
}

func (t *taskRepository) UpdateTask(taskRequest *models.Task) error {
	_, err := t.db.Update(t.taskTable).Set(
		goqu.Record{"title": taskRequest.Title, "description": taskRequest.Description, "status": taskRequest.Status},
	).Where(goqu.C("id").Eq(taskRequest.ID)).Executor().Exec()
	if err != nil {
		return merry.Wrap(err)
	}
	return nil
}

func (t *taskRepository) DeleteTask(taskID int64) error {
	if _, err := t.db.Delete(t.taskTable).Where(goqu.C("id").Eq(taskID)).Executor().Exec(); err != nil {
		return merry.Wrap(err)
	}
	return nil
}

func (t *taskRepository) AddTask(taskRequest *models.Task) (*models.Task, error) {
	taskRequest.CreatedAt = time.Now()
	r, err := t.db.Insert(t.taskTable).Rows(
		taskRequest,
	).Executor().Exec()

	if err != nil {
		return nil, merry.Wrap(err)
	}
	lastID, err := r.LastInsertId()
	if err != nil {
		return nil, merry.Wrap(err)
	}
	taskRequest.ID = lastID
	return taskRequest, nil
}

package controller

import (
	"net/http"
	"strconv"
	"todo/core/interactor"
	"todo/core/models"
	"todo/core/ports"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TaskControllerInterface interface {
	GetTaskList(c *gin.Context)
	GetTaskByID(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	AddTask(c *gin.Context)
}

type taskController struct {
	taskInteractor ports.TaskPortInterface
}

var TaskController TaskControllerInterface = &taskController{
	taskInteractor: interactor.TaskInteractor,
}

func (t *taskController) GetTaskList(c *gin.Context) {
	taskList, err := t.taskInteractor.GetAllTask()
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, taskList)
}

func (t *taskController) AddTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userID := c.GetFloat64("x-user-id")
	task.UserID = int64(userID)

	_, err := t.taskInteractor.AddTask(&task)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

func (t *taskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	task, err := t.taskInteractor.GetSingleTask(taskID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (t *taskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	task.ID = taskID
	if err := t.taskInteractor.UpdateTask(&task); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (t *taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if err := t.taskInteractor.DeleteTask(taskID); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

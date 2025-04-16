package handler

import (
	"net/http"
	"taskFlow/internal/schema"
	"taskFlow/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTasks(service *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := service.GetAllTasks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func GetTaskByID(service *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		task, err := service.GetTaskByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, task)
	}
}

func CreateTask(service *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task schema.Task

		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := task.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if task.ID == "" {
			task.ID = uuid.New().String()
		}

		if err := service.CreateTask(&task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, task)
	}
}

func UpdateTask(service *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		var task schema.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := task.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		existingTask, err := service.GetTaskByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		if task.Title != "" {
			existingTask.Title = task.Title
		}
		if task.Description != "" {
			existingTask.Description = task.Description
		}
		if task.Status != "" {
			existingTask.Status = task.Status
		}

		if err := service.UpdateTask(existingTask); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, existingTask)
	}
}

func DeleteTask(service *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		_, err := service.GetTaskByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		if err := service.DeleteTask(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"message": "Task deleted successfully"})
	}
}

package main

import (
	"log"
	"taskFlow/internal/config"
	"taskFlow/internal/handler"
	"taskFlow/internal/repository"
	"taskFlow/internal/schema"
	"taskFlow/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	dsn := cfg.ConnectionURL()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if err := db.AutoMigrate(&schema.Task{}); err != nil {
		log.Fatalf("Erro ao rodar migração: %v", err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)

	r := gin.Default()

	tasksGroup := r.Group("/tasks")
	{
		tasksGroup.GET("", handler.GetTasks(taskService))
		tasksGroup.GET("/:id", handler.GetTaskByID(taskService))
		tasksGroup.POST("", handler.CreateTask(taskService))
		tasksGroup.PUT("/:id", handler.UpdateTask(taskService))
		tasksGroup.DELETE("/:id", handler.DeleteTask(taskService))
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

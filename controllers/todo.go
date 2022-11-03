package controllers

import (
	"errors"
	"todo-app/database"
	"todo-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoRepo struct {
	Db *gorm.DB
}

func New() *TodoRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Todo{})
	return &TodoRepo{Db: db}
}

func (repository *TodoRepo) GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetTodos(repository.Db, &todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (repository *TodoRepo) GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	err := models.GetTodo(repository.Db, &todo, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (repository *TodoRepo) PostTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.PostTodo(repository.Db, &todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (repository *TodoRepo) UpdateTodo (c *gin.Context) {
	var todo models.Todo
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetTodo(repository.Db, &todo, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&todo)
	err = models.UpdateTodo(repository.Db, &todo, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (repository *TodoRepo) DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteTodo(repository.Db, &todo, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
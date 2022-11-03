package main

import (
	// "net/http"
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"todo-app/controllers"
)

// type todo struct {
// 	ID			string	`json: "id"`
// 	Activity	string	`json: "activity"`
// 	Status		bool	`json: "status"`
// }

// var todos = []todo{
// 	{ID: "1", Activity: "Meeting with client", Status: false},
// 	{ID: "2", Activity: "Create UI/UX Design", Status: false},
// 	{ID: "3", Activity: "Develop Backend Application", Status: false},
// }

// func getTodo(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, todos)
// }

// func postTodo(c *gin.Context) {
// 	var newTodo todo

// 	if err:= c.BindJSON(&newTodo); err != nil {
// 		return
// 	}

// 	todos = append(todos, newTodo)
// 	c.IndentedJSON(http.StatusCreated, newTodo)
// }



func main() {
	router := gin.Default()
	router.Use(cors.Default())

	todoRepo := controllers.New()
    router.GET("/todo", todoRepo.GetTodos)
    router.GET("/todo/:id", todoRepo.GetTodo)
    router.POST("/todo", todoRepo.PostTodo)
    router.PUT("/todo/:id", todoRepo.UpdateTodo)
    router.DELETE("/todo/:id", todoRepo.DeleteTodo)

    router.Run("localhost:8181")
}
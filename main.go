package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID					string			`json:"id"`
	Item 				string	 		`json:"item"`
	Completed 	bool 				`json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Belajar Golang", Completed: false},
	{ID: "2", Item: "Belajar RestAPI", Completed: false},
	{ID: "3", Item: "Belajar Gin", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context){
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated,newTodo)
}

func main() {
	Router := gin.Default()
	Router.GET("/todos", getTodos)
	Router.POST("/todos", addTodos)
	Router.Run("localhost:9090")
}
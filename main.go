package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
    ID        string `json:"id"`         // Unique id
    Item      string `json:"item"`       // Task desc
    Completed bool   `json:"completed"`  // Status of task 
}

//  Data
var todos = []Todo{
    {ID: "1", Item: "Clean Room", Completed: false }, 
    {ID: "2", Item: "Read Book", Completed: false }, 
    {ID: "3", Item: "Record Video", Completed: false }, 
}

// Diff struc that JSON
// Server -> Sending to Client convert the DS to JSON & vice versa 

// Endpoint for recive todo
func getTodos(context *gin.Context) { 
    context.IndentedJSON(http.StatusOK, todos)  
}

// Endpoint add to do -> Recieve data from client
func addTodo(context *gin.Context){
  var newTodo Todo

  // take whatever JSON inside request body & bind to new todo 
  if err := context.BindJSON(&newTodo); err != nil{
    return 
  }

  todos = append(todos, newTodo)

  context.IndentedJSON(http.StatusCreated, newTodo)
}

// iterate through the array and find specific todo
func getTodoById(id string) (*Todo, error) { 
  for i, t := range todos {
    if t.ID == id {
        return &todos[i], nil
    }
  }

  return nil, errors.New("todo not found")
}

// function for handler, extracting dynamic ID from the URL

func getTodo(context *gin.Context){

  id := context.Param("id")
  todo, err := getTodoById(id) 

  if err != nil {
      context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
      return
  }

  context.IndentedJSON(http.StatusOK, todo)
}


func toggleTodoStatus(context *gin.Context){

  id := context.Param("id")
  todo, err := getTodoById(id) 

  if err != nil {
      context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
      return
  }

  todo.Completed = !todo.Completed // flip boolean value


  context.IndentedJSON(http.StatusOK, todo)

}

// Server Creation & Endpoint
func main() {
  router := gin.Default()         // Create a new router
  router.GET("/todos", getTodos)
  router.GET("/todos/:id", getTodo)  // Endpoint Creation -> stored as id
  router.PATCH("/todos/:id", toggleTodoStatus)  // Endpoint Creation -> stored as id
  router.POST("/todos", addTodo)
  router.Run("localhost:9090")    
}

package main

import (
	"net/http"
  "errors",
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

// iteratre through the array and find specific todo
func getTodoById(id string)(*todo, error){
  for i, t:= range todos{
    if t.ID == id {
        return &todo[i], nil
    }
  }

  return nil, errors.New("todo not found")
}


// Server Creation & Endpoint
func main() {
  router := gin.Default()         // Create a new router
  router.GET("/todos", getTodos)  // Endpoint Creation
  router.POST("/todos", addTodo)
  router.Run("localhost:9090")    
}

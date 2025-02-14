package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Todo struct {
    ID        string `json:"id"`         // Unique id
    Item      string `json:"title"`      // Task desc
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

// incoming http request 
func getTodos(context *gin.Context) { 
    context.IndentedJSON(http.StatusOK, todos)  
}

// Server Creation & Endpoint
func main() {
  router := gin.Default()         // Create a new router
  router.GET("/todos", getTodos)  // Endpoint Creation
  router.Run("localhost:9090")    
}

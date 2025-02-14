package main

type todo struct {
    id string `json:"id"`
    item string `json:"title"`
    completed bool `json: "completed"`
}

var todos = []todo{
    {id: "1", item: "Clean Room", completed: false },
    {id: "2", item: "Read Book", completed: false },
    {id: "3", item: "Record Video", completed: false },
}

// Diff struc that JSON

// Client & the Server communicating through JSON
// Server -> Sending to Client convert the DS to JSON & vice versa 


// Server Creation

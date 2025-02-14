````md
# Todo API with Gin

This is a simple RESTful API for managing todos using the Gin framework in Go.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/todo-api-gin.git
   cd todo-api-gin
   ```
````

2. Install dependencies:
   ```sh
   go mod init todo-api
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
4. The API will be available at:
   ```
   http://localhost:9090
   ```

## API Endpoints

| Method  | Endpoint     | Description            |
| ------- | ------------ | ---------------------- |
| `GET`   | `/todos`     | Get all todos          |
| `GET`   | `/todos/:id` | Get a specific todo    |
| `POST`  | `/todos`     | Add a new todo         |
| `PATCH` | `/todos/:id` | Toggle todo completion |

## Usage

### GET /todos

Returns a list of all todos.

![GET Todos](Get.PNG)

### POST /todos

Adds a new todo.

![POST Todo](POST.PNG)

```

```

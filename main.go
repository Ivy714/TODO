package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` // e.g. "pending", "done"
}

var todos = []Todo{
	{ID: "1", Title: "寫 Go 作業", Status: "pending"},
	{ID: "2", Title: "看電腦美編", Status: "done"},
}

// GET 所有 TODO
func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// GET 單一 TODO
func getTodo(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// POST 新增 TODO
func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// PUT 更新 TODO
func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// DELETE 刪除 TODO
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func main() {
	r := gin.Default()

	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run() // 預設跑在 localhost:8080

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run(":8008")
}

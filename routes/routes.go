package routes

import (
	"database/sql"
	"simple-go-crud/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the router with routes
func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// API schema endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"endpoints": "CRUD operations for /users"})
	})

	// CRUD routes
	r.POST("/users", func(c *gin.Context) { controllers.CreateUser(c, db) })
	r.GET("/users/:id", func(c *gin.Context) { controllers.GetUser(c, db) })
	r.PUT("/users/:id", func(c *gin.Context) { controllers.UpdateUser(c, db) })
	r.DELETE("/users/:id", func(c *gin.Context) { controllers.DeleteUser(c, db) })
	r.GET("/users", func(c *gin.Context) { controllers.ListUsers(c, db) })

	return r
}

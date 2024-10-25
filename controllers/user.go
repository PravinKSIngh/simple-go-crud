package controllers

import (
	"database/sql"
	"net/http"
	"simple-go-crud/models"

	"github.com/gin-gonic/gin"
)

// The function `CreateUser` in Go handles creating a new user by parsing JSON data, inserting it into
// a database, and returning the created user with an assigned ID.
func CreateUser(c *gin.Context, db *sql.DB) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	smt, err := db.Prepare("INSERT INTO users(name, email) VALUES (?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to prepare SQL statement"})
		return
	}
	defer smt.Close()
	result, err := smt.Exec(user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to prepare SQL statement"})
		return
	}
	id, _ := result.LastInsertId()
	user.Id = int(id)
	c.JSON(http.StatusCreated, user)
}

// The function `GetUser` retrieves a user from a database based on the provided ID and returns the
// user details in JSON format.
func GetUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var user models.User
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// The function `UpdateUser` updates a user record in a database using data from a JSON request in a Go
// application.
func UpdateUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to prepare SQL statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to execute SQL statement"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// The DeleteUser function deletes a user from the database based on the provided ID.
func DeleteUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to prepare SQL statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to execute SQL statement"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// The ListUsers function retrieves user data from a database and returns it as JSON in a Gin context.
func ListUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to query users"})
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to scan user"})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

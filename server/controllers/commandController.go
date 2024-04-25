package controllers

import (
	"net/http"

	"server/db"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetCommands(c *gin.Context) {
	database := db.InitDB()
	defer database.Close()

	rows, _ := database.Query("SELECT id, command, status FROM commands")
	defer rows.Close()

	var commands []models.Command
	for rows.Next() {
		var cmd models.Command
		rows.Scan(&cmd.ID, &cmd.Command, &cmd.Status)
		commands = append(commands, cmd)
	}

	c.JSON(http.StatusOK, commands)
}

func AddCommand(c *gin.Context) {
	var newCmd models.Command
	if err := c.ShouldBindJSON(&newCmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.InitDB()
	defer database.Close()

	statement, _ := database.Prepare("INSERT INTO commands (command, status) VALUES (?, ?)")
	statement.Exec(newCmd.Command, "pending")

	c.JSON(http.StatusOK, gin.H{"message": "Command added"})
}

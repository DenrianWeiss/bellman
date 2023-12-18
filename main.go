package main

import (
	"github.com/DenrianWeiss/bellman/task"
	"github.com/DenrianWeiss/bellman/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	task.ScanJobCron()
	//task.ScanJob()
	// Start gin server
	r := gin.Default()
	// Add cors middleware
	r.Use(cors.Default())
	web.RegisterRoute(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

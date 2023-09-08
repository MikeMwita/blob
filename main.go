package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/api", func(c *gin.Context) {
		slackName := c.DefaultQuery("slack_name", "Michael")
		track := c.DefaultQuery("track", "backend")
		currentDay := time.Now().UTC().Format("Friday")
		currentUTC := time.Now().UTC()
		currentTime := currentUTC.Format("2006-01-02T15:04:05Z")
		githubFileURL := "https://github.com/MikeMwita/blob/blob/main/main/main.go"

		githubRepoURL := "https://github.com/MikeMwita/blob"

		responseData := gin.H{
			"slack_name":      slackName,
			"current_day":     currentDay,
			"utc_time":        currentTime,
			"track":           track,
			"github_file_url": githubFileURL,
			"github_repo_url": githubRepoURL,
			"status_code":     http.StatusOK,
		}

		c.JSON(http.StatusOK, responseData)
	})

	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS,GET,PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()

	}

}

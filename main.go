package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/send", sendFileHandler)

	router.Run()
}

func sendFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error :" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read file from request",
		})
		return
	}
	err = c.SaveUploadedFile(file, "./savedFile/"+file.Filename)
	if err != nil {
		log.Println("Error :" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save file from request",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{})
}

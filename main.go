package main

import (
	"net/http"
	"example.com/main/models"
	"github.com/gin-gonic/gin"
)



func main(){
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")

}
func getEvents(context *gin.Context){

	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event) //if the required fields are not there then there will be error and we are handling those errors by storing the error if generated
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"could not parse request"})
		return
	}

	event.ID = 1
	event.UserId = 1

	context.JSON(http.StatusCreated, gin.H{"message":"event created!", "event":event})

}
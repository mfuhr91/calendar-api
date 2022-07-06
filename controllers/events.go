package controllers

import (
	"calendar-api/models"
	"calendar-api/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type eventController struct{}

type EventController interface {
	Ping(c *gin.Context)
	Save(c *gin.Context)
	Delete(c *gin.Context)
}

var (
	eventRepo repositories.EventRepository
)

func NewEventController(repo repositories.EventRepository) *eventController {
	eventRepo = repo
	return &eventController{}
}

func (cont *eventController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (cont *eventController) GetAllEvents(c *gin.Context) {

	all, err := eventRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, all)
}

func (cont *eventController) Save(c *gin.Context) {

	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	saved, err := eventRepo.Save(event)
	if err != nil {
		log.Printf("error when saving the event - %s", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, saved)

}

func (cont *eventController) Update(c *gin.Context) {

	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	saved, err := eventRepo.Update(event)
	if err != nil {
		log.Printf("error when updating the event - %s", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, saved)

}

func (cont *eventController) Delete(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = eventRepo.Delete(event.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Msg": "deleted id: " + event.ID})

}

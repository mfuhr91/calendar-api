package router

import (
	"calendar-api/controllers"
	"calendar-api/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	eventRepository = repositories.NewEventRepository()
	cont            = controllers.NewEventController(eventRepository)
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(controllers.SetHeaders)

	api.GET("/ping", cont.Ping)
	api.GET("/events", cont.GetAllEvents)
	api.POST("/saveEvent", cont.Save)
	api.POST("/updateEvent", cont.Update)
	api.POST("/deleteEvent", cont.Delete)

	r.Static("static", "/static")
	r.StaticFile("manifest.json", "/manifest.json")
	r.StaticFile("calendar-1024.png", "/calendar-1024.png")
	r.LoadHTMLGlob("/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

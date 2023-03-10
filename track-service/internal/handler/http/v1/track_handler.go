package httphandler

import (
	"track-service/domain"

	"github.com/gin-gonic/gin"
)

func NewTrackHandler(trackUseCase domain.TrackUseCase) TrackHandler {
	return TrackHandler{TrackUseCase: trackUseCase}
}

type TrackHandler struct {
	TrackUseCase domain.TrackUseCase
}

func (handler *TrackHandler) Route(app *gin.Engine) {
	route := app.Group("/api/v1")
	route.POST("/track", handler.Add)
	route.GET("/track/find/:id", handler.GetById)
	route.GET("/track", handler.GetList)
	route.DELETE("/track/:id", handler.Delete)
	route.PUT("/track/:id", handler.Update)
}

func (handler *TrackHandler) Add(c *gin.Context) {
	var payload domain.TrackReq

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	result, err := handler.TrackUseCase.Add(&payload)

	if err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"message": "success",
		"data": map[string]any{"id": result},
	})	
}

func (handler *TrackHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	result, err := handler.TrackUseCase.GetById(id)

	if err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"message": "success",
		"data": result,
	})	
}

func (handler *TrackHandler) GetList(c *gin.Context) {
	result, err := handler.TrackUseCase.List()

	if err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"message": "success",
		"data": result,
	})	
}

func (handler *TrackHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := handler.TrackUseCase.Delete(id)

	if err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"message": "success",
	})	
}

func (handler *TrackHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var payload domain.TrackReq

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	result, err := handler.TrackUseCase.Update(id, &payload)

	if err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"message": "success",
		"data": result,
	})	
}
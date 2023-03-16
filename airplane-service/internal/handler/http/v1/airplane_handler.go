package httphandler

import (
	"airplane-service/domain"

	"github.com/gin-gonic/gin"
)

func NewAirplaneHandler(airplaneUseCase domain.AirplaneUseCase) AirplaneHandler {
	return AirplaneHandler{AirplaneUseCase: airplaneUseCase}
}

type AirplaneHandler struct {
	AirplaneUseCase domain.AirplaneUseCase
}

func (handler *AirplaneHandler) Route(app *gin.Engine) {
	route := app.Group("/api/v1")
	route.POST("/airplane", handler.Add)
	route.GET("/airplane/find/:id", handler.GetById)
	route.GET("/airplane", handler.GetList)
	route.DELETE("/airplane/:id", handler.Delete)
	route.PUT("/airplane/:id", handler.Update)
}

func (handler *AirplaneHandler) Add(c *gin.Context) {
	var payload domain.AirplaneReq

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	result, err := handler.AirplaneUseCase.Add(&payload)

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

func (handler *AirplaneHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	result, err := handler.AirplaneUseCase.GetById(id)

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

func (handler *AirplaneHandler) GetList(c *gin.Context) {
	result, err := handler.AirplaneUseCase.List()

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

func (handler *AirplaneHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := handler.AirplaneUseCase.Delete(id)

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

func (handler *AirplaneHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var payload domain.AirplaneReq

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	err := handler.AirplaneUseCase.Update(id, &payload)

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
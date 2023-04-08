package httphandler

import (
	"ticket-service/domain"

	"github.com/gin-gonic/gin"
)

func NewTicketHandler(ticketUseCase domain.TicketUseCase) TicketHandler {
	return TicketHandler{TicketUseCase: ticketUseCase}
}

type TicketHandler struct {
	TicketUseCase domain.TicketUseCase
}

func (handler *TicketHandler) Route(app *gin.Engine) {
	route := app.Group("/api/v1")
	route.POST("/ticket", handler.Add)
	route.DELETE("/ticket/:id", handler.Delete)
	route.PUT("/ticket/:id", handler.Update)
	route.GET("/ticket/find/:id", handler.GetById)
	route.GET("/ticket", handler.GetList)
}

func (handler *TicketHandler) Add(c *gin.Context) {
	var payload domain.TicketReq

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	result, err := handler.TicketUseCase.Add(&payload)

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

func (handler *TicketHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := handler.TicketUseCase.Delete(id)

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

func (handler *TicketHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var payload domain.TicketReq

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	err := handler.TicketUseCase.Update(id, &payload)

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

func (handler *TicketHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	result, err := handler.TicketUseCase.GetById(id)

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

func (handler *TicketHandler) GetList(c *gin.Context) {
	result, err := handler.TicketUseCase.List()

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

func (handler *TicketHandler) Search(c *gin.Context) {
	id := c.Query("arrival")

	result, err := handler.TicketUseCase.GetById(id)

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

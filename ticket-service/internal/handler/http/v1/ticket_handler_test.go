package httphandler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"ticket-service/domain"
	httphandler "ticket-service/internal/handler/http/v1"
	mockusecase "ticket-service/internal/mocks/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTicketHandler_Add(t *testing.T) {

	gin.SetMode(gin.TestMode)

	mockTicketUC := new(mockusecase.TicketUseCase)

	var idTicket string = "0f10cff9-4b5a-4e5a-b976-56bf56efa280"

	mockPayload := domain.TicketReq{
		TrackId: "d414381b-62fc-451f-a00f-b5a16995dc12",
		AirplaneId: "788e7859-379f-42f2-88f3-41fccad0e234",
		Datetime: "2022-12-21T13:00:00Z",
		Price: 1000000,
	}

	payload, err := json.Marshal(mockPayload)
	
	if err != nil {
		t.Errorf("Json marshal payload error: %s", err)
	}

	t.Run("success", func(t *testing.T) {
		mockTicketUC.On("Add", mock.AnythingOfType("*domain.TicketReq")).Return(idTicket, nil)

		rr := httptest.NewRecorder()

		g := gin.Default()
		httphandler.NewTicketHandler(
			mockTicketUC,
		)

		ticketHttpHandler := httphandler.NewTicketHandler(mockTicketUC)

		ticketHttpHandler.Route(g)

		request, err := http.NewRequest(http.MethodPost, "/api/v1/ticket", strings.NewReader(string(payload)))
		assert.NoError(t, err)

		g.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"status": "ok",
			"message": "success",
			"data": map[string]any{"id": idTicket},
		})

		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockTicketUC.AssertExpectations(t)


	})
}
package controller

import (
	"errors"
	"example/httpserver/config"
	"example/httpserver/models"
	"example/httpserver/response"
	"example/httpserver/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Config       *config.Configurations
	EventService *services.EventService
}

func NewEventHandler() *Event {
	return &Event{}
}

func (evt *Event) PostEventInfoHandler(ctx *gin.Context) {
	resp := new(response.Response)
	event := new(models.Event)
	if err := ctx.BindJSON(&event); err != nil {
		log.Default().Printf("[Error] [PostEventInfoHandler] [BindJson] : %s", err)
		err = errors.New("invalid request params. Please check the input payload")
		resp.SendError(ctx, http.StatusBadRequest, err)
		return
	}
	evt.EventService = new(services.EventService)
	eventChan := make(chan models.Event)
	outputChan := make(chan interface{})
	evt.EventService.Config = *evt.Config
	go evt.EventService.EventWorkerService(ctx, eventChan, outputChan)
	eventChan <- *event
	select {
	case data := <-outputChan:
		eventOut, ok := data.(models.StrcturedEvent)
		if !ok {
			log.Default().Printf("[Error] [PostEventInfoHandler] [chanelout] : %s", data)
			resp.SendError(ctx, http.StatusInternalServerError, errors.New("coudn't process the data"))
			return
		}
		resp.SendData(ctx, map[string]interface{}{
			"data": eventOut,
		})
	case <-time.After(5 * time.Second):
		resp.SendError(ctx, http.StatusRequestTimeout, errors.New("timeout waiting for response"))
		return
	}

}

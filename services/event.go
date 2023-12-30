package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"example/httpserver/config"
	"example/httpserver/models"
	"example/httpserver/utils"
	"io"
	"log"
	"net/http"
	"time"
)

type EventService struct {
	Config config.Configurations
}

func (eventSrv *EventService) EventWorkerService(ctx context.Context, eventChan chan models.Event, out chan<- interface{}) {
	eventInput := <-eventChan
	outEvent := utils.CopyToStructuredEventStruct(&eventInput)
	payload, err := json.Marshal(outEvent)
	if err != nil {
		out <- errors.New("caoudnt marshal the json")
		return
	}
	req, err := http.NewRequest(http.MethodPost, eventSrv.Config.WebhokUrl, bytes.NewBuffer(payload))
	if err != nil {
		out <- errors.New("caoudnt create the http request to webhook - make sure the url in config file is correct")
		return
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(30*time.Second))
	defer cancel()
	req = req.WithContext(ctx)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		out <- errors.New("caoudnt send the http request to webhook - make sure the url in config file is correct")
		return
	}
	defer res.Body.Close()
	httpresp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Default().Printf("[Error] [EventWorkerService] [httpresponse] [read] : %s", err)
	}
	log.Default().Printf("[Error] [EventWorkerService] [httpresponse] : %s", httpresp)
	out <- outEvent
}

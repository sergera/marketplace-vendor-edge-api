package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sergera/marketplace-vendor-edge-api/internal/conf"
	"github.com/sergera/marketplace-vendor-edge-api/internal/domain"
)

type WorkerService struct {
	host        string
	port        string
	contentType string
	client      *http.Client
}

func NewWorkerService() *WorkerService {
	conf := conf.GetConf()
	return &WorkerService{
		conf.WorkerHost,
		conf.WorkerPort,
		"application/json; charset=UTF-8",
		&http.Client{},
	}
}

func (s WorkerService) Post(route string, jsonData []byte) error {
	request, err := http.NewRequest("POST", s.host+":"+s.port+"/"+route, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("failed to create post request: " + err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := s.client.Do(request)
	if err != nil {
		log.Println("failed to perform worker post request: " + err.Error())
		return err
	}

	defer response.Body.Close()
	return nil
}

func (s WorkerService) Put(route string, jsonData []byte) error {
	request, err := http.NewRequest("PUT", s.host+":"+s.port+"/"+route, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("failed to create put request: " + err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := s.client.Do(request)
	if err != nil {
		log.Println("failed to perform worker put request: " + err.Error())
		return err
	}

	defer response.Body.Close()
	return nil
}

func (s WorkerService) UpdateStatus(o domain.OrderModel) error {
	m, err := json.Marshal(o)
	if err != nil {
		log.Println("Failed to marshal event model into json")
		return err
	}

	err = s.Put("update-order-status", m)
	if err != nil {
		return err
	}

	return nil
}

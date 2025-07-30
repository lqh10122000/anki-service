package service

import (
	"bytes"
	"complaint-service/internal/model"
	"complaint-service/internal/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeskService interface {
	InputDesk() error
	GetAllDesk() ([]model.Desks, error)
}

type deskService struct {
	repo repository.DeskRepository
}

func NewDeskService(repo repository.DeskRepository) DeskService {
	return &deskService{repo: repo}
}

func (s *deskService) InputDesk() error {
	url := "http://127.0.0.1:8765"
	payload := map[string]interface{}{
		"action":  "deckNames",
		"version": 6,
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)

	var result struct {
		Result []string `json:"result"`
		Error  string   `json:"error"`
	}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil
	}

	for i := 0; i < len(result.Result); i++ {
		foundDesk, err := s.repo.FindAllByName(result.Result[i])
		if err != nil && len(foundDesk) == 0 {
			s.repo.SaveDesk(result.Result[i])
		}
		fmt.Println("full deskkk: ", result.Result[i])
	}

	return nil
}

func (s *deskService) GetAllDesk() ([]model.Desks, error) {
	return s.repo.FindAll()
}

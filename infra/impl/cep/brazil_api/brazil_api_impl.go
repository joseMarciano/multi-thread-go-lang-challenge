package brazil_api

import (
	"encoding/json"
	"multithread-challenge/model"
	"net/http"
	"strings"
)

const BASE_URL = "https://brasilapi.com.br/api/cep/v1/{cep}"

type resposeBrazilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (responseBrazil *resposeBrazilApi) toModel() *model.ZipCodeInfo {
	return &model.ZipCodeInfo{
		Code:   responseBrazil.Cep,
		State:  responseBrazil.State,
		City:   responseBrazil.City,
		Origin: "brazil_api",
	}
}

type BrazilApiImpl struct {
	httpClient *http.Client
}

func NewBrazilApi() *BrazilApiImpl {
	return &BrazilApiImpl{
		httpClient: http.DefaultClient,
	}
}

func (b *BrazilApiImpl) GetZipCodeInfo(zip string) (*model.ZipCodeInfo, error) {
	response, err := http.Get(strings.ReplaceAll(BASE_URL, "{cep}", zip))

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responseBody resposeBrazilApi
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	if err != nil {
		return nil, err
	}

	return responseBody.toModel(), nil
}

package via_cep

import (
	"encoding/json"
	"multithread-challenge/model"
	"net/http"
	"strings"
)

const BASE_URL = "https://viacep.com.br/ws/{cep}/json/"

type responseViaCorreio struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (responseBrazil *responseViaCorreio) toModel() *model.ZipCodeInfo {
	return &model.ZipCodeInfo{
		Code:   responseBrazil.Cep,
		State:  responseBrazil.Uf,
		City:   responseBrazil.Localidade,
		Origin: "via_cep_api",
	}
}

type ViaCepApiImpl struct {
	httpClient *http.Client
}

func NewViaCepApi() *ViaCepApiImpl {
	return &ViaCepApiImpl{
		httpClient: http.DefaultClient,
	}
}

func (b *ViaCepApiImpl) GetZipCodeInfo(zip string) (*model.ZipCodeInfo, error) {
	response, err := http.Get(strings.ReplaceAll(BASE_URL, "{cep}", zip))

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responseBody responseViaCorreio
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	if err != nil {
		return nil, err
	}

	return responseBody.toModel(), nil
}

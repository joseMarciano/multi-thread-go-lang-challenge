package main

import (
	"encoding/json"
	"fmt"
	"multithread-challenge/infra/impl/cep/brazil_api"
	"multithread-challenge/infra/impl/cep/composition"
	"multithread-challenge/infra/impl/cep/via_cep"
	"multithread-challenge/model/interfaces"
	"os"
)

func main() {
	var compositionZipCodeApi = composition.NewZipCodeCompositionApi([]interfaces.ZipCodeApi{
		brazil_api.NewBrazilApi(),
		via_cep.NewViaCepApi(),
	})

	cep := os.Args[1:][0]

	var zipCodeResponse, err = compositionZipCodeApi.GetZipCodeInfo(cep)
	if err != nil {
		fmt.Println("Error on get brazil zip code", err)
	}

	err = json.NewEncoder(os.Stdout).Encode(zipCodeResponse)
	if err != nil {
		panic(err)
	}

}

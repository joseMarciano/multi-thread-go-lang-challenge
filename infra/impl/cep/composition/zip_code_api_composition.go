package composition

import (
	"fmt"
	"multithread-challenge/model"
	"multithread-challenge/model/interfaces"
	"time"
)

type ZipCodeCompositionApi struct {
	zipCodeApis []interfaces.ZipCodeApi
	channel     chan *model.ZipCodeInfo
}

func NewZipCodeCompositionApi(zipCodeApis []interfaces.ZipCodeApi) *ZipCodeCompositionApi {
	return &ZipCodeCompositionApi{
		zipCodeApis: zipCodeApis,
		channel:     make(chan *model.ZipCodeInfo),
	}
}

func (z *ZipCodeCompositionApi) GetZipCodeInfo(zipCode string) (*model.ZipCodeInfo, error) {
	//var channel = make(chan *model.ZipCodeInfo)

	for _, api := range z.zipCodeApis {
		go z.callZipCode(zipCode, api)
	}

	select {
	case info := <-z.channel:
		return info, nil
	case <-time.After(time.Second):
		return nil, fmt.Errorf("error on get zip code %s timeout", zipCode)
	}

}

func (z *ZipCodeCompositionApi) callZipCode(zipCode string, zipCodeApi interfaces.ZipCodeApi) {
	info, err := zipCodeApi.GetZipCodeInfo(zipCode)
	if err != nil {
		panic(err)
	}

	z.channel <- info
}

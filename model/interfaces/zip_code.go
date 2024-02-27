package interfaces

import "multithread-challenge/model"

type ZipCodeApi interface {
	GetZipCodeInfo(zipCode string) (*model.ZipCodeInfo, error)
}

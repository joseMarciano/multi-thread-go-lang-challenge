package model

type ZipCodeInfo struct {
	Code   string `json:"code"`
	City   string `json:"city"`
	State  string `json:"state"`
	Origin string `json:"origin"`
}

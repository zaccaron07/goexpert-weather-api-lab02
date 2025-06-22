package entity

import (
	"errors"
	"regexp"
)

var ErrInvalidZipcode = errors.New("invalid zipcode")

type ZipcodeRequest struct {
	Cep string `json:"cep"`
}

func NewZipcode(cep string) (*ZipcodeRequest, error) {
	matched, _ := regexp.MatchString(`^\d{8}$`, cep)
	if !matched {
		return nil, ErrInvalidZipcode
	}
	return &ZipcodeRequest{Cep: cep}, nil
}

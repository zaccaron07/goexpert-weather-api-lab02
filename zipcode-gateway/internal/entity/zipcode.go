package entity

import (
	"errors"
	"regexp"
)

var errInvalidZipcode = errors.New("invalid zipcode")

type ZipcodeRequest struct {
	Cep string
}

func NewZipcode(cep string) (*ZipcodeRequest, error) {
	matched, _ := regexp.MatchString(`^\d{8}$`, cep)
	if !matched {
		return nil, errInvalidZipcode
	}
	return &ZipcodeRequest{Cep: cep}, nil
}

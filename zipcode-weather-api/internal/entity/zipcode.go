package entity

import (
	"errors"
	"regexp"
)

type Zipcode struct {
	CEP        string
	Bairro     string
	Localidade string
	UF         string
}

func NewZipcode(zipcode string) (*Zipcode, error) {
	if !isValidZipcode(zipcode) {
		return nil, errors.New("invalid zipcode")
	}
	return &Zipcode{
		CEP: zipcode,
	}, nil
}

func isValidZipcode(zipcode string) bool {
	match, _ := regexp.MatchString("^[0-9]{8}$", zipcode)
	return match
}

package usecase

import (
	"context"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/entity"
)

type ZipcodeInput struct {
	CEP string
}

type ZipcodeOutput struct {
	CEP        string `json:"cep"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

type GetZipcodeUseCase struct {
	ZipcodeRepository entity.ZipcodeRepositoryInterface
}

func NewGetZipcodeUseCase(zipcodeRepository entity.ZipcodeRepositoryInterface) *GetZipcodeUseCase {
	return &GetZipcodeUseCase{
		ZipcodeRepository: zipcodeRepository,
	}
}

func (c *GetZipcodeUseCase) Execute(ctx context.Context, input ZipcodeInput) (ZipcodeOutput, error) {
	zipcode, err := entity.NewZipcode(input.CEP)
	if err != nil {
		return ZipcodeOutput{}, err
	}

	zipcodeDto, err := c.ZipcodeRepository.Get(ctx, zipcode.CEP)
	if err != nil {
		return ZipcodeOutput{}, err
	}

	zipecodeOutput := ZipcodeOutput{
		CEP:        zipcodeDto.CEP,
		Bairro:     zipcodeDto.Bairro,
		Localidade: zipcodeDto.Localidade,
		UF:         zipcodeDto.UF,
	}

	return zipecodeOutput, nil
}

package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/zaccaron07/goexpert-weather-api-lab02/zipcode-weather-api/internal/entity"
)

type ZipcodeRepository struct{}

func NewZipcodeRepository() *ZipcodeRepository {
	return &ZipcodeRepository{}
}

func (r *ZipcodeRepository) Get(zipcodeAddress string) (entity.Zipcode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://viacep.com.br/ws/%s/json", zipcodeAddress), nil)
	if err != nil {
		log.Printf("Failed to initialize request: %v", err)
		return entity.Zipcode{}, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Request failed: %v", err)
		return entity.Zipcode{}, err
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading the response: %v", err)
		return entity.Zipcode{}, err
	}

	var zipcode entity.Zipcode
	err = json.Unmarshal(resp, &zipcode)
	if err != nil {
		log.Printf("Error parsing response: %v", err)
		return entity.Zipcode{}, err
	}
	return zipcode, nil
}

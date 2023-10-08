package service

import (
	"math/rand"

	"github.com/adenhidayatuloh/assignment3-011/dto"
	"github.com/adenhidayatuloh/assignment3-011/entity"
	"github.com/adenhidayatuloh/assignment3-011/repo/database"
)

// func pada service untuk upadate data pada database, dan akan mengembalikan response dari package dto
func UpdatePeriod() (*dto.Response, error) {

	randomWater := rand.Intn(100) + 1
	randomWind := rand.Intn(100) + 1

	var payload = entity.WaterWind{
		Water: randomWater,
		Wind:  randomWind,
	}

	var Response = dto.Response{
		Water: randomWater,
		Wind:  randomWind,
	}

	err := database.UpdateWaterWind(payload)

	if err != nil {
		return nil, err
	}
	return &Response, nil

}

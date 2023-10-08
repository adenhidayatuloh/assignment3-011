package service

import (
	"assigment3-GLNG-08-011/dto"
	"assigment3-GLNG-08-011/entity"
	"assigment3-GLNG-08-011/repo/database"
	"math/rand"
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

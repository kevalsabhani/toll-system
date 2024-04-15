package obu

import (
	"math"
	"math/rand"
)

type OBUData struct {
	OBUId int     `json:"obuId"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func NewOBUData() *OBUData {
	lat, long := generateLatLong()
	return &OBUData{
		OBUId: generateOBUId(),
		Lat:   lat,
		Long:  long,
	}
}

func generateOBUId() int {
	return rand.Intn(math.MaxInt)
}

func generateLatLong() (float64, float64) {
	return generateCoordinate(), generateCoordinate()
}

func generateCoordinate() float64 {
	return float64(rand.Intn(100)) + rand.Float64()
}

package obu

import (
	"math"
	"math/rand"
)

// OBUData defines OBU propetires
type OBUData struct {
	OBUId int     `json:"obuId"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

// NewOBUData returns an object of OBUData
func NewOBUData() *OBUData {
	lat, long := generateLatLong()
	return &OBUData{
		OBUId: generateOBUId(),
		Lat:   lat,
		Long:  long,
	}
}

// generateOBUId generates random OBU id
func generateOBUId() int {
	return rand.Intn(math.MaxInt)
}

// generateLatLong generate random latitude, longitude
func generateLatLong() (float64, float64) {
	return generateCoordinate(), generateCoordinate()
}

// generateCoordinate random coordinate
func generateCoordinate() float64 {
	return float64(rand.Intn(100)) + rand.Float64()
}

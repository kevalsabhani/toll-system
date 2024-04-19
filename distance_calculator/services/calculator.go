package services

import (
	"math"

	"github.com/kevalsabhani/toll-calculator/client/services"
)

var coords = [2][2]float64{}

func DistanceCal(data services.OBUData) float64 {
	if coords[1][0] == 0 && coords[1][1] == 0 {
		coords[1][0] = data.Lat
		coords[1][1] = data.Long
		return 0
	}
	coords[0][0] = coords[1][0]
	coords[0][1] = coords[1][1]
	coords[1][0] = data.Lat
	coords[1][1] = data.Long
	return math.Sqrt(math.Pow((coords[1][0]-coords[0][0]), 2) + math.Pow((coords[1][1]-coords[0][1]), 2))
}

package types

type Distance struct {
	Value     float64 `json:"value"`
	OBUId     int     `json:"obuId"`
	Timestamp int64   `json:"timestamp"`
}

type Invoice struct {
	OBUId         int     `json:"obuId"`
	TotalDistance float64 `json:"totalDistance"`
	Amount        float64 `json:"amount"`
}

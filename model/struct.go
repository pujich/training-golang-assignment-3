package model

type M map[string]interface{}

type WaterWind struct {
	Water int `json:"water" form:"water"`
	Wind  int `json:"wind" form:"wind"`
}

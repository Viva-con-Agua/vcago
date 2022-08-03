package vmod

//Position represents the lat and lng coordinates for latitude and longitude.
type Position struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

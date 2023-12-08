package types

type Distance struct {
	Value float64 `json:"value"`
	UID   int     `json:"uid"`
	Unix  int64   `json:"unix"`
}

type GPSData struct {
	UID int     `json:"uid"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"long"`
}

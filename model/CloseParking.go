package model

type FareInfoResponse struct {
	RespCode    string  `json:"respCode"`
	RespMessage string  `json:"respMessage"`
	VehicleNo   string  `json:"vehicleNo"`
	EntryDt     string  `json:"entryDt"`
	ParkingId   int     `json:"parkingId"`
	ParkingFare float64 `json:"parkingFare"`
}

package model

import (
	"fmt"
	"math/rand"
	"time"
)

type CompoundResult struct {
	CouncilID        string    `json:"COUNCILID"`
	SearchType       string    `json:"SEARCHTYPE"`
	SearchValue      string    `json:"SEARCHVALUE"`
	RefNo            string    `json:"REFNO"`
	TransDate        time.Time `json:"TRANSDATE"`
	KodHasil         string    `json:"KODHASIL"`
	AmtToPay         float64   `json:"AMNTOPAY"`
	StatusKmp        string    `json:"STATUSKMP"`
	Location         string    `json:"LOCATION"`
	Offence          string    `json:"OFFENCE"`
	VehicleNo        string    `json:"VEHICLENO"`
	Error            *string   `json:"ERROR"`
	StatusRayuan     string    `json:"STATUSRAYUAN"`
	RayuanTandaHarga float64   `json:"RAYUANTANDAHARGA"`
}

func GenerateCompoundResult() CompoundResult {
	rand.Seed(time.Now().UnixNano())

	return CompoundResult{
		CouncilID:        "MBPJ",
		SearchType:       "VEHICLENO",
		SearchValue:      "BMB" + randomDigits(4),
		RefNo:            "CMP" + randomAlphaNum(10),
		TransDate:        time.Now().AddDate(0, 0, -rand.Intn(30)), // dalam 30 hari ke belakang
		KodHasil:         "524",
		AmtToPay:         float64(rand.Intn(100)+50) + 0.5,
		StatusKmp:        "UNPAID",
		Location:         "Jalan SS" + fmt.Sprint(rand.Intn(20)+1),
		Offence:          "Parking at No Parking Zone",
		VehicleNo:        "BMB" + randomDigits(4),
		StatusRayuan:     "PENDING",
		RayuanTandaHarga: 100.00,
	}
}

func randomAlphaNum(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func randomDigits(length int) string {
	const digits = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}

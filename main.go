package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gomock/model"
	"log"
	"math/rand"
	"strings"
	"time"
)

type KompaunRequest struct {
	Carian        string `json:"Carian"`
	CarianKompaun string `json:"Carian_Kompaun"`
	Searchvalue   string `json:"searchvalue"`
}

type KompaunDetail struct {
	NoKompaun     string `json:"nokompaun"`
	NoDaftar      string `json:"nodaftar"`
	Status        string `json:"status"`
	TarikhKompaun string `json:"KP_TarikhKompaun"`
	Jumlah        int    `json:"jumlah"`
	Masa          string `json:"masa"`
	IC            string `json:"ic"`
	StatusKompaun int    `json:"KP_StatusKompaun"`
	CukaiJalan    string `json:"cukaijalan"`
	NamaKesalahan string `json:"SAL_NAMA"`
}

type KompaunResponse struct {
	Status string `json:"status"`
}

type ParkmaxCompoundResponse struct {
	Status string `json:"Status"`
}

type UpdateCukaiResponse struct {
	Message string `json:"message"`
}

func main() {
	compounds := map[string]int{
		"KMP100001": 50,
		"KMP100002": 30,
		"KMP100003": 75,
		"KMP100004": 60,
		"KMP100005": 90,
		"KMP100006": 45,
		"KMP100007": 100,
		"KMP100008": 35,
		"KMP100009": 25,
		"KMP100010": 80,
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Post("/mphtj_online/Carian_api.cfm", func(c *fiber.Ctx) error {
		returnEmpty := false
		request := KompaunRequest{
			Carian:        c.FormValue("Carian"),
			CarianKompaun: c.FormValue("Carian_Kompaun"),
		}

		log.Println(request)
		log.Println(request.CarianKompaun)
		log.Println(request.Carian)

		if request.CarianKompaun != "" {
			compoundAmnt := compounds[request.CarianKompaun]
			if compoundAmnt == 0 {
				return c.SendStatus(fiber.StatusNotFound)
			}

			listResponse := []KompaunDetail{}
			data := KompaunDetail{
				NoKompaun:     request.CarianKompaun,
				NoDaftar:      "",
				Status:        "Kompaun",
				TarikhKompaun: "",
				Jumlah:        compoundAmnt,
				Masa:          "February, 23 2024 16: 01: 14 +0800",
				IC:            "",
				StatusKompaun: 5,
				CukaiJalan:    "",
				NamaKesalahan: "TELAH MELETAK KERETA MOTOR ATAU MENUNGGU DIMANA-MANA JALAN SELAIN DARI TEMPAT LETAKKERETA",
			}
			listResponse = append(listResponse, data)

			return c.JSON(listResponse)
		}

		list := []KompaunDetail{}
		i := 0
		for k, v := range compounds {
			list = append(list, KompaunDetail{
				NoKompaun:     k,
				NoDaftar:      "",
				Status:        "Kompaun",
				TarikhKompaun: "",
				Jumlah:        v,
				Masa:          "February, 23 2024 16: 01: 14 +0800",
				IC:            "",
				StatusKompaun: 5,
				CukaiJalan:    "",
				NamaKesalahan: "TELAH MELETAK KERETA MOTOR ATAU MENUNGGU DIMANA-MANA JALAN SELAIN DARI TEMPAT LETAKKERETA",
			})
			i++
			if i == 2 {
				break
			}
		}

		if strings.HasSuffix(request.CarianKompaun, "EMPTY") {
			returnEmpty = true
		}

		if !returnEmpty {
			return c.JSON(list)
		} else {
			return c.JSON([]KompaunDetail{})
		}

	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber jalan bangsat 💥")
	})

	app.Post("/splk/sonic.php", func(c *fiber.Ctx) error {
		response := KompaunResponse{
			Status: "BERJAYA",
		}
		return c.JSON(response)
	})

	app.Post("/api/updatePayCukai", func(c *fiber.Ctx) error {
		response := UpdateCukaiResponse{
			Message: "success",
		}
		return c.JSON(response)
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Hello " + name + " 👊")
	})

	app.Post("/UpdatePaymentWS", func(c *fiber.Ctx) error {
		response := ParkmaxCompoundResponse{
			Status: "success",
		}
		return c.JSON(response)
	})

	app.Post("/GetPenyataProcess", func(c *fiber.Ctx) error {
		response := map[string]string{
			"NOPENYATA": "P000000",
		}
		return c.JSON(response)
	})

	app.Post("/GetCompDetailToPaySP", func(c *fiber.Ctx) error {
		list := []model.CompoundResult{}

		request := KompaunRequest{
			Searchvalue: c.FormValue("searchvalue"),
		}

		response := model.GenerateCompoundResult()

		if request.Searchvalue == "error" {
			err := "Error"
			response.Error = &err
		}

		if request.Searchvalue == "0" {
			return c.JSON([]model.CompoundResult{})
		}

		list = append(list, response)
		return c.JSON(list)
	})

	app.Get("/commercepay/transaction", func(c *fiber.Ctx) error {
		resp := model.Response{
			Resutl: model.Result{
				TransactionNumber:         "TX-" + randString(12),
				ReferenceCode:             "ORX0000668" + randString(2),
				Status:                    1, // fixed 0
				CurrencyCode:              randomCurrency(),
				Amount:                    int64(rand.Intn(9) + 10),
				ChannelID:                 rand.Intn(10) + 1,
				ProviderTransactionNumber: "PTX-" + randString(10),
				CreationTime:              time.Now().UTC(),
				Remark:                    "Mocked response " + randString(5),
				ProviderChannelID:         "CH-" + randString(6),
				ProviderPaymentMethod:     randomPayMethod(),
			},
		}
		return c.JSON(resp)
	})

	app.Post("/terminal/vehicle/fare/info", func(c *fiber.Ctx) error {
		data := []byte(`{"respCode": "0000",
		"respMessage": "Success",
		"vehicleNo": "ABC55431",
		"entryDt": "2025-03-24 09:35:18",
		"parkingId": 46,
		"parkingFare": 1.5
	}`)

		var resp model.FareInfoResponse
		if err := json.Unmarshal(data, &resp); err != nil {
			panic(err)
		}

		return c.JSON(resp)

	})

	err := app.Listen(":8090")
	if err != nil {
		panic(err)
	}
}

func randString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomCurrency() string {
	currencies := []string{"IDR", "MYR", "USD", "SGD", "EUR"}
	return currencies[rand.Intn(len(currencies))]
}

func randomPayMethod() string {
	methods := []string{"CARD", "BANK_TRANSFER", "EWALLET", "QRIS", "VIRTUAL_ACCOUNT"}
	return methods[rand.Intn(len(methods))]
}

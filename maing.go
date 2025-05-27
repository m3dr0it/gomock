package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type KompaunRequest struct {
	Carian        string `json:"Carian"`
	CarianKompaun string `json:"Carian_Kompaun"`
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
		var request KompaunRequest
		if err := c.BodyParser(&request); err != nil {
			err.Error()
		}

		if request.CarianKompaun != "" {
			compoundAmnt := compounds[request.CarianKompaun]
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

			return c.JSON(data)
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
		return c.JSON(list)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber jalan bangsat 💥")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Hello " + name + " 👊")
	})

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

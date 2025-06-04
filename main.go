package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

// gw update dari branch1
// gw update dari branch_2
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

type KompaunResponse struct {
	Status string `json:"status"`
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
		return c.JSON(list)
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

	err := app.Listen(":8090")
	if err != nil {
		panic(err)
	}
}

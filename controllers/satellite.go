package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/federicolivarez/challengeMeli/utils"
)

type Satellite struct {
	Name string `json:"name"`
	Distance float64 `json:"distance"`
	Message []string `json:"message"`
}

type SatelliteReq struct {
	Satellites []Satellite `Satellite:"satellites"`
}

func GetLocation(c *fiber.Ctx) error {
	satelliteReq := new(SatelliteReq)
    err := c.BodyParser(satelliteReq)

    if err != nil {
        fmt.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success":  false,
            "message": "No se puede calcular la posicion con la informacion recibida",
        })
    }

	Positions := utils.GetLocation( satelliteReq.Satellites[0].Distance,satelliteReq.Satellites[1].Distance,satelliteReq.Satellites[2].Distance)

	MessageCompleted := utils.GetMessage(satelliteReq.Satellites[0].Message,satelliteReq.Satellites[1].Message,satelliteReq.Satellites[2].Message)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"position": fiber.Map{
			"x": Positions.Ejex,
			"y": Positions.Ejey,
		},
		"message": MessageCompleted,
	})
}


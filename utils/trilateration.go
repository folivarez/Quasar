package utils

import (
    "math"
	"strings"
	"github.com/federicolivarez/challengeMeli/config"
)

type Positions struct {
	Ejex     float64   `json:"ejex"`
	Ejey     float64  `json:"ejey"`
}

func GetLocation(distancea,distanceb,distancec float64 )(*Positions) {

	configuration := config.GetConfiguration()

	var xa float64 = configuration.Satellites.Satellite_1_x
	var ya float64 = configuration.Satellites.Satellite_1_y
	var xb float64 = configuration.Satellites.Satellite_2_x
	var yb float64 = configuration.Satellites.Satellite_2_y
	var xc float64 = configuration.Satellites.Satellite_3_x
	var yc float64 = configuration.Satellites.Satellite_3_y

	var S float64 = (math.Pow(xc, 2.) - math.Pow(xb, 2.) + math.Pow(yc, 2.) - math.Pow(yb, 2.) + math.Pow(distanceb, 2.) - math.Pow(distancec, 2.)) / 2.0
	var T float64 = (math.Pow(xa, 2.) - math.Pow(xb, 2.) + math.Pow(ya, 2.) - math.Pow(yb, 2.) + math.Pow(distanceb, 2.) - math.Pow(distancea, 2.)) / 2.0
	var y float64 = ((T * (xb - xc)) - (S * (xb - xa))) / (((ya - yb) * (xb - xc)) - ((yc - yb) * (xb - xa)))
	var x float64 = ((y * (ya - yb)) - T) / (xb - xa)

	p := Positions{Ejex: x, Ejey: y }
    return &p
}


func GetMessage(message1, message2, message3 []string)(string){

	var mensajeCompleto [5]string

	for i := 0; i < len(mensajeCompleto); i++ {
		if message1[i] != ""{
			mensajeCompleto[i] = message1[i]		
		}
		if message2[i] != ""{
			mensajeCompleto[i] = message2[i]		
		}
		if message3[i] != ""{
			mensajeCompleto[i] = message3[i]		
		}
	}

	return strings.Join(mensajeCompleto[:], " ")
}


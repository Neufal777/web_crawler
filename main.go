package main

import (
	"github.com/draco/core"
)

func main() {

	// fuente := core.Source{
	// 	Id:       core.ID,
	// 	Country:  core.COUNTRY,
	// 	Vertical: core.VERTICAL,
	// 	Url:      core.URL,

	// 	Category: []core.Category{
	// 		{
	// 			Type: core.FOR_SALE,
	// 			Urls: []string{
	// 				"https://www.pisos.com/venta/pisos-a_coruna",
	// 				"https://www.pisos.com/venta/pisos-barcelona",
	// 				"https://www.pisos.com/venta/pisos-girona",
	// 			},
	// 		},
	// 		{
	// 			Type: core.FOR_RENT,
	// 			Urls: []string{
	// 				"https://www.pisos.com/alquiler/pisos-a_coruna",
	// 				"https://www.pisos.com/alquiler/pisos-barcelona",
	// 			},
	// 		},
	// 	},
	// }

	prueba := []string{
		"https://www.pisos.com/alquiler/pisos-a_coruna",
	}

	core.GetAdUrls(prueba)
}

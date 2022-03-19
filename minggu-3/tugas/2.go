package main

import (
	"fmt"
	"math"
)

func main() {
	var massa_tubuh int
	var w_merkurius, w_venus, w_bumi, w_mars, w_yupiter, w_saturnus, w_uranus, w_neptunus float64

	fmt.Scan(&massa_tubuh)

	w_merkurius = float64(massa_tubuh) * (float64(0.38) * float64(9.8))
	w_venus = float64(massa_tubuh) * (float64(0.91) * float64(9.8))
	w_bumi = float64(massa_tubuh) * float64(9.8)
	w_mars = float64(massa_tubuh) * (float64(0.38) * float64(9.8))
	w_yupiter = float64(massa_tubuh) * (float64(2.37) * float64(9.8))
	w_saturnus = float64(massa_tubuh) * (float64(0.92) * float64(9.8))
	w_uranus = float64(massa_tubuh) * (float64(0.89) * float64(9.8))
	w_neptunus = float64(massa_tubuh) * (float64(1.13) * float64(9.8))

	fmt.Println(math.Round(w_merkurius), math.Round(w_venus), math.Round(w_bumi), math.Round(w_mars), math.Round(w_yupiter), math.Round(w_saturnus), math.Round(w_uranus), math.Round(w_neptunus))

}

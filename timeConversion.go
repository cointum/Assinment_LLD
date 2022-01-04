package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {

	mil:=30099458424

	y := (mil/1000)
	// convert int to string 
	q := strconv.Itoa(y)
	var r string = "s"

	t := q + r

	u, _ := time.ParseDuration(t)
	//conver hour into days
	var totalDays float64 = (u.Hours()) / 24

	fmt.Printf("Milliseconds to christmas: %d\n", mil)
	fmt.Println("Seconds to christmas: ", y)

	fmt.Printf("Minutes to christmas: %.f\n", math.Round(u.Minutes()))

	fmt.Printf("Hours to christmas: %.f\n", math.Round(u.Hours()))

	fmt.Printf("Days  to christmas: %.f\n", math.Round(totalDays))
}

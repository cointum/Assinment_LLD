package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	y := 30099458
	q := strconv.Itoa(y)
	var r string = "s"
	t := q + r

	u, _ := time.ParseDuration(t)
	var totalDays float64 = (u.Hours()) / 24

	fmt.Printf("Milliseconds to christmas: %d\n", u.Milliseconds())
	fmt.Println("Seconds to christmas: ", y)

	fmt.Printf("Minutes to christmas: %.f\n", math.Round(u.Minutes()))

	fmt.Printf("Hours to christmas: %.f\n", math.Round(u.Hours()))

	fmt.Printf("Days  to christmas: %.f\n", math.Round(totalDays))
}

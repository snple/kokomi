package util

import (
	"math/rand"
	"strconv"
)

func Round(f float64, n int) float64 {
	formatNum, _ := strconv.ParseFloat(strconv.FormatFloat(f, 'f', n, 64), 64)
	return formatNum
}

func RandNum(min, max float64) float64 {
	return rand.Float64()/((1-0)/(max-min)) + min
}

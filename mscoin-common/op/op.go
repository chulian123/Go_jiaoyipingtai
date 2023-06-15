package op

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func DivN(x float64, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x/y), 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return parseFloat
}
func MulN(x float64, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x*y), 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return parseFloat
}

func AddN(x float64, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x+y), 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return parseFloat
}
func ReduceN(x float64, y float64, n int) float64 {
	s := fmt.Sprintf("%d", n)
	parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%."+s+"f", x-y), 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return parseFloat
}

// FloorFloat 1.245  保留两位小数点 1.25  金融方面的应用 1.24  2.0/3.0
func FloorFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Floor(val*ratio) / ratio
}

// RoundFloat 进位运算  1.245  保留两位小数点 1.25  金融方面的应用 1.24 而不是 1.25
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

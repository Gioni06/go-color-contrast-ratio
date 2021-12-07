package colorcontrast

import (
	"image/color"
	"math"
	"sort"
)

func reverse(input []float64) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func luminance(r, g, b float64) float64 {
	var val [3]float64
	val[0] = r
	val[1] = g
	val[2] = b

	for i, colorChannel := range val {
		if colorChannel < .03928 {
			val[i] = colorChannel / 12.92
		} else {
			val[i] = math.Pow((colorChannel+0.055)/1.055, 2.4)
		}
	}

	return (0.2126 * val[0]) + (0.7152 * val[1]) + (0.0722 * val[2])
}

func toArithmetic(col color.Color) (r, g, b float64) {
	ri, gi, bi, _ := col.RGBA()
	r = float64(ri) / math.MaxUint16
	g = float64(gi) / math.MaxUint16
	b = float64(bi) / math.MaxUint16
	return
}

func Calculate(a, b color.Color) float64 {
	luminancePair := []float64{
		luminance(toArithmetic(a)),
		luminance(toArithmetic(b)),
	}
	sort.Float64s(luminancePair)
	reverse(luminancePair)
	return (luminancePair[0] + 0.05) / (luminancePair[1] + 0.05)
}

package testing

import (
	"encoding/csv"
	"fmt"
	. "github.com/gioni06/go-color-contrast"
	"image/color"
	"io"
	"log"
	"os"
	"regexp"
	"testing"
)

func TrimTrailingZeros(numberString string) string {
	rgx := regexp.MustCompile(`(\.[0-9]*[1-9])0+$|\.0*$`)
	return rgx.ReplaceAllString(numberString, "$1")
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	return
}

type testColor struct {
	ContrastRatio string
	Rating        string
	ColorA        color.Color
	ColorB        color.Color
	ColorAHex     string
	ColorBHex     string
}

func TestContrastRatio(t *testing.T) {
	var testColors []testColor
	csvFile, _ := os.Open("./All-DB.csv")
	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		colorA, err := ParseHexColor("#" + record[2])
		colorB, err := ParseHexColor("#" + record[3])

		if err != nil {
			continue
		}
		testColors = append(testColors, testColor{
			ContrastRatio: record[0],
			Rating:        record[1],
			ColorA:        colorA,
			ColorB:        colorB,
			ColorAHex:     "#" + record[2],
			ColorBHex:     "#" + record[3],
		})
	}

	for _, tc := range testColors {
		rawContrastRatioString := fmt.Sprintf("%f", ContrastRatioGo(tc.ColorA, tc.ColorB))
		input := rawContrastRatioString[0:4]
		input = TrimTrailingZeros(input)
		test := TrimTrailingZeros(tc.ContrastRatio)

		if input != test {
			t.Errorf("contrast ratio not matched for colors %s, %s. Given %s, Actual %s", tc.ColorAHex, tc.ColorBHex, test, input)
		}
	}
}

func TestName(t *testing.T) {
	colorA := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	colorB := color.RGBA{
		R: 231,
		G: 48,
		B: 55,
		A: 255,
	}
	fmt.Println(ContrastRatioGo(colorA, colorB))
}

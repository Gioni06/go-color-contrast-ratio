# go-color-contrast-ratio

Calculate the color contrast ratio specified in [WCAG 2.1](https://www.w3.org/TR/WCAG/#contrast-minimum)

```go
package main

import (
	"fmt"
	"github.com/Gioni06/go-color-contrast-ratio"
	"image/color"
)

func main()  {
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
	// 4.307044780466779
	fmt.Println(colorcontrast.Calculate(colorA, colorB))
}
```
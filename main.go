package main

import (
	"fmt"
	"os"
)

func main() {
	imageWidth := 256
	imageHeight := 256

	f, err := os.Create("output.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))
	if err != nil {
		fmt.Println(err)
		return
	}

	for j := 0; j < imageHeight; j++ {
		fmt.Printf("\rScanlines remaining: %d", imageHeight-j)
		for i := 0; i < imageWidth; i++ {
			pixelColor := Color{
				e: [3]float64{
					float64(i) / float64(imageWidth-1),  // r
					float64(j) / float64(imageHeight-1), // g
					0.0,                                 // b
				},
			}

			_, err := f.WriteString(WriteColor(pixelColor))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println("\rDone.                         ")
}

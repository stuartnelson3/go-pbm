package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	f, _ := os.Open("./out.pbm")
	defer f.Close()

	reader := bufio.NewReader(f)
	ln, _, _ := reader.ReadLine()
	if string(ln) != "P1" {
		fmt.Println("Not a PBM file: ", ln)
		return
	}
	var width, height int
	fmt.Fscanln(reader, &width, &height)

	r := image.Rect(0, 0, width, height)
	gray := image.NewGray16(r)

	// fmt.Println(width)
	// fmt.Println(height)

	for i := 0; i < height; i++ {
		ln, _, _ = reader.ReadLine()
		for j, n := range ln {
			c := color.White
			if string(n) == "1" {
				c = color.Black
			}
			gray.Set(j, i, c)
		}

	}

	// Write image
	f, _ = os.Create("pbm.png")
	png.Encode(f, gray.SubImage(gray.Rect))
}

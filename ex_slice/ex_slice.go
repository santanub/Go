package main

import(
	"tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)

	for x:=0; x < dy; x++ {
		picture[x] = make([]uint8, dx)

		for y:=0; y < dx; y++ {
			picture[x][y] = uint8(y^x)
		}
	}

	return picture
}

func main() {
	pic.Show(Pic)
}

package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func testPos(path string) {
	dst_map := robotgo.OpenBitmap(path)
	defer robotgo.FreeBitmap(dst_map)

	x, y := robotgo.FindBitmap(dst_map)
	fmt.Printf("stone x %d y %d", x, y)
	x, y = robotgo.GetMousePos()
	fmt.Printf("now x %d y %d", x, y)
}

package main

import (
	"fmt"
	"image"
	"os"

	"github.com/labstack/gommon/log"
)

// whilescreenbase 多次尋找圖片
// fullimg 圖片路徑
// count 尋找幾次
// args[0] 匹配程度
func whilescreenbase(fullimg string, count int, args ...float64) (succ bool, x int, y int) {
	matchNumber := 0.01
	if len(args) == 1 {
		matchNumber = args[0]
	}

	//確認有圖案
	file, _ := os.Open(fullimg)
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Errorf("load pngName Error:%s, pngName: %s", err, pngName)
		file.Close()
		return
	}

	file.Close()

	imgh := c.Height
	imgw := c.Width
	copyFileContents(pngName, "tmp.png")
	bit_map := robotgo.OpenBitmap("tmp.png")

	defer robotgo.FreeBitmap(bit_map)

	count := 20
	if len(jcount) == 1 {
		count = int(jcount[0])
	}
	log.Info("screen", pngName)
	log.Info("count", count)
	for {
		robotgo.Sleep(1)

		fx, fy := robotgo.FindBitmap(bit_map)

		fmt.Println("FindBitmap------", fx, fy)
		if fx != -1 && fy != -1 {
			succ = true
			x = fx + (imgw / 2)
			y = fy + (imgh / 2)
			return
		}
		count = count - 1
		if count == 0 {
			succ = false
			x = -1
			y = -1
			return
		}
	}

}

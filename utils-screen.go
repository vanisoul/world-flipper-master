package main

import (
	"image"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

var basewhilecount = 5

// findscreen 多次尋找圖片
// fullimg 圖片路徑
// args[0] 匹配程度 defult 0.01
// =>
// 找到與否
func findscreen(fullimg string, args ...float64) bool {
	matchNumber := 0.01
	if len(args) == 1 {
		matchNumber = args[0]
	}
	succ, _, _ := whilescreen(fullimg, matchNumber)
	return succ
}

// whilescreen 多次尋找圖片
// fullimg 圖片路徑
// args[0] 匹配程度 defult 0.01
// =>
// 找到與否 以及 x,y
func whilescreen(fullimg string, args ...float64) (succ bool, x int, y int) {
	matchNumber := 0.01
	if len(args) == 1 {
		matchNumber = args[0]
	}
	succ, x, y = whilescreenbase(fullimg, basewhilecount, matchNumber)
	return
}

// whilescreenbase 多次尋找圖片
// fullimg 圖片路徑
// count 尋找幾次
// args[0] 匹配程度
// =>
// 找到與否 以及 x,y
func whilescreenbase(fullimg string, count int, matchNumber float64) (succ bool, x int, y int) {

	//確認有圖案
	file, _ := os.Open(fullimg)
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Errorf("load pngName Error:%s, pngName: %s", err, fullimg)
		file.Close()
		return
	}
	file.Close()

	imgh := c.Height
	imgw := c.Width

	bit_map := robotgo.OpenBitmap(fullimg)

	defer robotgo.FreeBitmap(bit_map)

	log.Infof("whilescreenbase, fullimg: %s", fullimg)
	log.Infof("whilescreenbase, count: %d", count)
	log.Infof("whilescreenbase, matchNumber: %d", matchNumber)
	for {
		tmp_bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
		defer robotgo.FreeBitmap(tmp_bitmap)
		fx, fy := robotgo.FindBitmap(bit_map, tmp_bitmap, matchNumber)

		if fx != -1 && fy != -1 {
			succ = true
			x = fx + (imgw / 2)
			y = fy + (imgh / 2)
			log.Infof("whilescreenbase-Success, x:%d, y:%d", fx, fy)
			return
		}
		count = count - 1
		if count == 0 {
			succ = false
			x = -1
			y = -1
			log.Infof("whilescreenbase-error")
			return
		}
	}

}

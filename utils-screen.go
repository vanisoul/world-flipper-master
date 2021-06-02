package main

import (
	"fmt"
	"image"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

var basewhilecount = 20

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
	count_tmp := 0
	for count_tmp < count {
		tmp_bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
		defer robotgo.FreeBitmap(tmp_bitmap)
		fx, fy := robotgo.FindBitmap(bit_map, tmp_bitmap, matchNumber)
		if fx != -1 && fy != -1 {
			succ = true
			x = fx + (imgw / 2) + infox
			y = fy + (imgh / 2) + infoy
			log.Infof("whilescreenbase-Success, count_tmp: %d, x:%d, y:%d", count_tmp, x, y)
			setLog("whilescreenbase", "尋找圖片成功", fmt.Sprintf("img: %s, count_tmp:%d, x:%d, y:%d", fullimg, count_tmp, x, y))
			return
		}
		log.Infof("whilescreenbase-wait, count_tmp: %d", count_tmp)
		setLog("whilescreenbase", "尋找圖片中", fmt.Sprintf("img: %s, count_tmp:%d", fullimg, count_tmp))
		robotgo.Sleep(1)
		count_tmp = count_tmp + 1
		if count_tmp == count {
			succ = false
			x = -1
			y = -1
			log.Infof("whilescreenbase-error")
			setLog("whilescreenbase", "尋找圖片失敗", fmt.Sprintf("img: %s, count_tmp:%d", fullimg, count_tmp))
			return
		}
	}
	return
}

//儲存定義的的視窗 |
// args 儲存至 log/${args[0]}/${agrs[1]}.....
func savescreen(args ...string) int {

	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	imgName := fmt.Sprintf("log/%d.png", r)

	if len(args) > 0 {
		pathstr := path.Join(args...)
		imgName = fmt.Sprintf("log/%s/%d.png", pathstr, r)
		err := os.MkdirAll(path.Dir(imgName), os.ModePerm)
		if err != nil {
			log.Error(err)
		}
	}

	bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
	defer robotgo.FreeBitmap(bitmap)
	robotgo.SaveBitmap(bitmap, imgName)

	return r
}

//看到指定圖片 就截圖存放
func checkImgSaveImg(fullImg string, count int, args ...float64) bool {
	matchNumber := 0.01
	if len(args) == 1 {
		matchNumber = args[0]
	}
	succ, _, _ := whilescreenbase(fullImg, count, matchNumber)
	if succ {
		savescreen(thisID)
		return true
	}
	savescreen(thisID)
	return false
}

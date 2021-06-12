package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

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

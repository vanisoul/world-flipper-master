package main

import "github.com/go-vgo/robotgo"

func leftMouseClick() {
	robotgo.MouseClick("left")
}

func rightMouseClick() {
	robotgo.MouseClick("right")
}

//左鍵點擊圖片正中間
//fullimg 圖片路徑
//args[0] 匹配程度
func leftMouseClickImg(fullimg string, args ...float64) bool {
	matchNumber := 0.01
	if len(args) == 1 {
		matchNumber = args[0]
	}

	succ, x, y := whilescreen(fullimg, matchNumber)
	if succ {
		robotgo.MoveMouse(x, y)
		leftMouseClick()
		return true
	}
	return false
}

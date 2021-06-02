package main

import "github.com/go-vgo/robotgo"

func leftMouseClick() {
	robotgo.MouseClick("left")
}

func rightMouseClick() {
	robotgo.MouseClick("right")
}

//左鍵點擊多圖片的其中一張正中間 |
//fullImgs 多圖片路徑 |
//matchNumber 匹配程度
func leftMouseClickImgMany(matchNumber float64, fullImgs ...string) bool {

	for _, fullimg := range fullImgs {
		succ, x, y := whilescreenbase(fullimg, 2, matchNumber)
		if succ {
			robotgo.MoveMouse(x, y)
			leftMouseClick()
			return true
		}
	}
	return false
}

//左鍵點擊圖片正中間 |
//fullimg 圖片路徑 |
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

//找到圖後 點選另一張圖 |
// findImg : 如果看到這張圖 |
// clickImg : 點選這張圖 |
// args[0] : find匹配程度 |
// args[1] : click匹配程度
func checkImgClickOtherImg(findImg string, clickImg string, args ...interface{}) bool {
	matchNumberFind := 0.01
	matchNumberClick := 0.01
	if len(args) >= 1 && args[0] != nil {
		matchNumberFind, _ = args[0].(float64)
	}
	if len(args) >= 1 && args[1] != nil {
		matchNumberClick, _ = args[1].(float64)
	}

	skipGameSucc := findscreen(findImg, matchNumberFind)
	if skipGameSucc {
		skipOKSucc := leftMouseClickImg(clickImg, matchNumberClick)
		if !skipOKSucc {
			return false
		}
	} else {
		return false
	}
	return true
}

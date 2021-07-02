package main

import (
	"github.com/go-vgo/robotgo"
)

var yDifficulty = 0
var yBoss = 0
var yEvery = 0
var imgBoss = "remaining.pngOrstroke.png"
var imgDifficulty = "updateList.pngOritemExchange.png"
var choeseBossSeq = 0       //這次選擇的關卡
var choeseDifficultySeq = 0 //這次選擇的關卡
var checkpointConfig CheckpointConfig
var tmpRFrequency = 0
var tmpAuto = ""      //紀錄auto狀態 如果是auto模式才有用到
var choseAuto = false //紀錄這次有沒有被改變狀態

func main() {
	//初始化ID
	infoID()

	tmpcheckpointConfig, _ := LoadCheckpointConfig()
	adbinit(tmpcheckpointConfig.Nox)
	tmpcheckpointConfig.Type = "" //故意改變讓一開始進入回主選單

	for {
		//如果config有變動 需要重新回到主頁
		checkpointConfig, _ = LoadCheckpointConfig()
		if tmpcheckpointConfig != checkpointConfig || choseAuto {
			tmpcheckpointConfig = checkpointConfig
			if tmpAuto != "repalay" {
				tmpAuto = checkpointConfig.Type
			}
			reStartImg := []string{}
			reStartFunc := []func(x int, y int){}
			reStartImg, reStartFunc = addStartRaising(reStartImg, reStartFunc)
			reStartImg, reStartFunc = addRgoGame3(reStartImg, reStartFunc)

			cmainImg := []string{}
			cmainFunc := []func(x int, y int){}
			cmainImg, cmainFunc = addMain2(cmainImg, cmainFunc)

			haveOneImgsExecFunc(1, 0.05, false, reStartImg, reStartFunc...)
			haveOneImgsExecFunc(10, 0.1, false, cmainImg, cmainFunc...)

			robotgo.Sleep(8)
			choseAuto = false
		}

		//主流程
		mainImg := []string{}
		mainFunc := []func(x int, y int){}
		mainImg, mainFunc = addGameLogs(mainImg, mainFunc)
		mainImg, mainFunc = addJoinMain(mainImg, mainFunc)
		mainImg, mainFunc = addMainMission(mainImg, mainFunc)
		mainImg, mainFunc = addImgBoss(mainImg, mainFunc)
		mainImg, mainFunc = addImgDifficulty(mainImg, mainFunc)
		mainImg, mainFunc = addMainMission(mainImg, mainFunc)

		//重復關卡
		freeImg := []string{}
		freeFunc := []func(x int, y int){}
		freeImg, freeFunc = addGreat(freeImg, freeFunc)
		freeImg, freeFunc = addMultiplayer(freeImg, freeFunc)
		freeImg, freeFunc = addYES(freeImg, freeFunc)
		freeImg, freeFunc = addOK(freeImg, freeFunc)
		freeImg, freeFunc = addNotRecruit(freeImg, freeFunc)
		freeImg, freeFunc = addGoPaly(freeImg, freeFunc)
		freeImg, freeFunc = addStop(freeImg, freeFunc)
		freeImg, freeFunc = addGoGame3(freeImg, freeFunc)
		freeImg, freeFunc = addNext1(freeImg, freeFunc)
		freeImg, freeFunc = addRePlay(freeImg, freeFunc)
		freeImg, freeFunc = addExitYes(freeImg, freeFunc)

		//較少用到
		dayImg := []string{}
		dayFunc := []func(x int, y int){}
		dayImg, dayFunc = addDayGift(dayImg, dayFunc)
		dayImg, dayFunc = addDayClose(dayImg, dayFunc)
		dayImg, dayFunc = addNetworkerrorOK(dayImg, dayFunc)
		dayImg, dayFunc = addGmaeOver(dayImg, dayFunc)
		dayImg, dayFunc = addExitHalfway(dayImg, dayFunc)

		haveOneImgsExecFunc(1, 0.05, false, mainImg, mainFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, mainImg, mainFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, freeImg, freeFunc...)
		haveOneImgsExecFunc(1, 0.05, false, dayImg, dayFunc...)

	}
}

func choeseBoss(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yBoss + seq*yEvery
	if seq < 6 {
		mouseClick(x, y)
		robotgo.Sleep(3)
	} else {
		ys := y_tmp + yBoss + 4*yEvery
		ye := y_tmp + yBoss + 3*yEvery
		AdbShellInputSwipe(x, ys, x, ye)
		choeseBoss(seq - 1)
		robotgo.Sleep(1)
	}
}

// func fchoeseBoss(seq int) {
// 	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
// 	y := y_tmp + 460 + seq*175
// 	if seq < 6 {
// 		mouseClick(x, y)
// 		robotgo.Sleep(3)
// 	} else {
// 		ys := y_tmp + 460 + 4*175
// 		ye := y_tmp + 460 + 3*175
// 		AdbShellInputSwipe(x, ys, x, ye)
// 		choeseBoss(seq - 1)
// 		robotgo.Sleep(1)
// 	}
// }

// func rchoeseBoss(seq int) {
// 	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
// 	y := y_tmp + 310 + seq*178
// 	if seq < 7 {
// 		mouseClick(x, y)
// 		robotgo.Sleep(3)
// 	} else {
// 		ys := y_tmp + 310 + 4*178
// 		ye := y_tmp + 310 + 3*178
// 		AdbShellInputSwipe(x, ys, x, ye)
// 		choeseBoss(seq - 1)
// 		robotgo.Sleep(1)
// 	}
// }

func choeseDifficulty(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yDifficulty + seq*175
	if seq < 6 {
		mouseClick(x, y)
		robotgo.Sleep(3)
	} else {
		ys := y_tmp + yDifficulty + 2*175
		ye := y_tmp + yDifficulty + 1*175
		AdbShellInputSwipe(x, ys, x, ye)
		choeseDifficulty(seq - 1)
		robotgo.Sleep(1)
	}
}

// func bfchoeseDifficulty(seq int) {
// 	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
// 	y := y_tmp + 310 + seq*175
// 	if seq < 6 {
// 		mouseClick(x, y)
// 		robotgo.Sleep(3)
// 	} else {
// 		ys := y_tmp + 310 + 2*175
// 		ye := y_tmp + 310 + 1*175
// 		AdbShellInputSwipe(x, ys, x, ye)
// 		choeseDifficulty(seq - 1)
// 		robotgo.Sleep(1)
// 	}
// }

// func acchoeseDifficulty(seq int) {
// 	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
// 	y := y_tmp + 460 + seq*175
// 	if seq < 6 {
// 		mouseClick(x, y)
// 		robotgo.Sleep(3)
// 	} else {
// 		ys := y_tmp + 460 + 2*175
// 		ye := y_tmp + 460 + 1*175
// 		AdbShellInputSwipe(x, ys, x, ye)
// 		choeseDifficulty(seq - 1)
// 		robotgo.Sleep(1)
// 	}
// }

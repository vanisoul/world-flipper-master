package main

import (
	"strconv"

	"github.com/go-vgo/robotgo"
)

var yDifficulty = 0
var yBoss = 0
var yEvery = 0
var imgBoss = "defultImg.png"
var imgDifficulty = "defultImg.png"
var choeseBossSeq = 0       //這次選擇的關卡
var choeseDifficultySeq = 0 //這次選擇的關卡
var checkpointConfig CheckpointConfig
var tmpRFrequency = 0
var choseAuto = false //紀錄這次有沒有被改變狀態
var status = 0
var notthink = 0 //多少次迴圈沒動作 1000次沒動作 就關閉重啟

func main() {
	//初始化ID
	infoID()

	tmpcheckpointConfig, _ := LoadCheckpointConfig()
	tmpcheckpointConfig.Type = "" //故意改變讓一開始進入回主選單
	for {
		//如果config有變動 需要重新回到主頁
		checkpointConfig, _ = LoadCheckpointConfig()
		if tmpcheckpointConfig != checkpointConfig || choseAuto {
			tmpcheckpointConfig = checkpointConfig
			adbinit(checkpointConfig.Nox)
			status = 0
		}

		//重開遊戲
		if status == 0 {
			//關閉遊戲
			closeApp()
			robotgo.Sleep(2)
			//啟動遊戲
			openApp()

			status = 1
		}

		// 到首頁
		if status == 1 {
			toMainImg := []string{}
			toMainFunc := []func(x int, y int){}
			toMainImg, toMainFunc = addJoinMain(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addOK(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addExitHalfway(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addDownloadOK(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addMainMissionMainOK(toMainImg, toMainFunc)
			haveOneImgsExecFunc(1, 0.05, false, toMainImg, toMainFunc...)
		}

		//確認已到首頁
		if status == 2 {
			//判斷體力是否滿
			dowhatImg := []string{}
			dowhatFunc := []func(x int, y int){}
			//如果體力滿就狀態5 消耗體力
			if checkpointConfig.RFeatures {
				dowhatImg, dowhatFunc = addFullOfEnergyMain(dowhatImg, dowhatFunc)
			}
			dowhatImg, dowhatFunc = addNotFullOfEnergyMain(dowhatImg, dowhatFunc)

			haveOneImgsExecFunc(1, 0.05, false, dowhatImg, dowhatFunc...)
		}

		//消耗體力
		if status == 5 {
			runOnlyPlayImg := []string{}
			runOnlyPlayFunc := []func(x int, y int){}
			runOnlyPlayImg, runOnlyPlayFunc = addJoinActivity(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addImgBoss(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addImgDifficulty(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addYES(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addGoGameOnly(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addGoGameMaze(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addGmaeOver(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addGameOverOK(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addNext1(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addLvUP(runOnlyPlayImg, runOnlyPlayFunc)
			runOnlyPlayImg, runOnlyPlayFunc = addRePlay(runOnlyPlayImg, runOnlyPlayFunc)
			haveOneImgsExecFunc(1, 0.05, false, runOnlyPlayImg, runOnlyPlayFunc...)
		}

		//進入免費房間
		if status == 6 {
			runActivityImgImg := []string{}
			runActivityImgFunc := []func(x int, y int){}
			if checkpointConfig.RFeatures {
				runActivityImgImg, runActivityImgFunc = addFullOfEnergy(runActivityImgImg, runActivityImgFunc)
			}
			if checkpointConfig.Type == "activity" {
				runActivityImgImg, runActivityImgFunc = addJoinActivity(runActivityImgImg, runActivityImgFunc)
			} else if checkpointConfig.Type == "boss" {
				runActivityImgImg, runActivityImgFunc = addJoinBoss(runActivityImgImg, runActivityImgFunc)
				runActivityImgImg, runActivityImgFunc = addJoinBoss2(runActivityImgImg, runActivityImgFunc)
			}

			runActivityImgImg, runActivityImgFunc = addImgBoss(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addImgDifficulty(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addMultiplayer(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addYES(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addNotRecruit(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addOnlyEachOther(runActivityImgImg, runActivityImgFunc)
			haveOneImgsExecFunc(1, 0.05, false, runActivityImgImg, runActivityImgFunc...)
		}

		//開始免費房間流程
		if status == 66 {
			runFreeRoomImg := []string{}
			runFreeRoomFunc := []func(x int, y int){}

			runFreeRoomImg, runFreeRoomFunc = addGoPaly(runFreeRoomImg, runFreeRoomFunc)
			runFreeRoomImg, runFreeRoomFunc = addGreat(runFreeRoomImg, runFreeRoomFunc)
			runFreeRoomImg, runFreeRoomFunc = addExitOK(runFreeRoomImg, runFreeRoomFunc)
			runFreeRoomImg, runFreeRoomFunc = addStop(runFreeRoomImg, runFreeRoomFunc)
			haveOneImgsExecFunc(1, 0.05, false, runFreeRoomImg, runFreeRoomFunc...)
		}

		if notthink > checkpointConfig.Notthink {
			savescreen("notthink", strconv.Itoa(status))
			status = 0
			notthink = 0
		}
		notthink++
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

package main

import (
	"github.com/go-vgo/robotgo"
)

var yDifficulty = 0
var yBoss = 0

func main() {

	//初始化設定 未來要判斷視窗位置
	infoScreen(-40, -40, 600, 1050)
	//初始化ID
	infoID()
	tmpDifficulty := 0
	tmpNumber := 0
	tmpType := ""
	for {
		//如果config有變動 需要重新回到主頁
		checkpointConfig, _ := LoadCheckpointConfig()
		if tmpDifficulty != checkpointConfig.Difficulty || tmpNumber != checkpointConfig.Number || tmpType != checkpointConfig.Type {
			tmpDifficulty = checkpointConfig.Difficulty
			tmpNumber = checkpointConfig.Number
			tmpType = checkpointConfig.Type
			haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("startRaising.png")},
				func(x, y int) {
					haveOneImgsLeft(1, 0.05, true, getSystemImg("return.png"))
					haveOneImgsLeft(1, 0.05, true, getSystemImg("disband.png"))
				})
			haveOneImgsLeft(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"))
		}

		//開啟遊戲
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("gameLogo.png")}, func(x, y int) {
			leftMouseClick(x, y)
		})
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("joinMain.png")}, func(x, y int) {
			leftMouseClick(x, y)
		})

		//進入bose戰鬥or活動戰鬥
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("mainMission.png")}, func(x, y int) {
			if checkpointConfig.Type == "boss" {
				haveOneImgsLeft(10, 0.05, true, getSystemImg("joinBoss.png"))
				yDifficulty = 200
				yBoss = 310
			} else if checkpointConfig.Type == "activity" {
				haveOneImgsLeft(10, 0.05, true, getSystemImg("joinActivity.png"))
				yDifficulty = 310
				yBoss = 200
			}
		})

		//處理選擇boss關卡
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg("stroke.png")}, func(x, y int) { choeseBoss(checkpointConfig.Number) })

		//處理選擇難度
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg("itemExchange.png")}, func(x, y int) { choeseDifficulty(checkpointConfig.Difficulty) })

		//多人遊玩
		haveOneImgsLeft(1, 0.05, true, getSystemImg("multiplayer.png"))
		//使用強化點數
		haveOneImgsLeft(1, 0.05, true, getSystemImg("YES.png"))

		//閒置
		haveOneImgsLeft(1, 0.05, true, getSystemImg("OK.png"))

		//如果有需要領取獎勵 關閉活動通知
		haveOneImgsLeft(1, 0.05, true, getSystemImg("dayGift.png"))
		haveOneImgsLeft(1, 0.05, true, getSystemImg("dayClose.png"))

		//如果沒有招募
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg("notRecruit.png")}, func(x, y int) {
			haveOneImgsLeft(5, 0.01, false, getSystemImg("startRaising.png"))
			haveOneImgsLeft(1, 0.01, false, getSystemImg("YrandomRecruitment.png"))
			haveOneImgsLeft(1, 0.01, false, getSystemImg("YfollowersPublic.png"))
			haveOneImgsLeft(5, 0.01, false, getSystemImg("goRecruit.png"))
		})

		//讚的圖示 變成隨機招募
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg("great.png")}, func(x, y int) {
			haveOneImgsLeft(5, 0.01, false, getSystemImg("startRaising.png"))
			haveOneImgsLeft(5, 0.01, false, getSystemImg("NrandomRecruitment.png"))
			haveOneImgsLeft(5, 0.01, false, getSystemImg("NfollowersPublic.png"))
			haveOneImgsLeft(5, 0.01, false, getSystemImg("goRecruit.png"))
		})

		//GO的圖示 開始
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg("goPaly.png")}, func(x, y int) {
			haveOneImgsLeft(5, 0.01, false, getSystemImg("goGame.png"))
		})

		//偵測到戰鬥中跳出
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg("stop.png")}, func(x, y int) {
			leftMouseClick(x, y)
			haveOneImgsLeft(5, 0.01, false, getSystemImg("exit.png"))
			haveOneImgsLeft(5, 0.01, false, getSystemImg("exitYes.png"))
		})

		//有錯誤問題
		haveOneImgsLeft(1, 0.01, false, getSystemImg("exitHalfway.png"))
		haveOneImgsLeft(1, 0.01, false, getSystemImg("errorOK.png"))

		robotgo.Sleep(3)
	}
}

func choeseBoss(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yBoss + seq*110
	if seq < 6 {
		leftMouseClick(x, y)
	} else {
		ys := y_tmp + 310 + 4*110
		ye := y_tmp + 310 + 1*110
		robotgo.MoveMouse(x, ys)
		robotgo.MouseToggle(`down`, `left`)
		robotgo.Sleep(1)
		robotgo.MoveMouse(x, ye)
		robotgo.Sleep(1)
		robotgo.MouseToggle(`up`, `left`)
		choeseBoss(seq - 3)
		robotgo.Sleep(1)
	}
}

func choeseDifficulty(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yDifficulty + seq*110
	if seq < 6 {
		leftMouseClick(x, y)
	} else {
		ys := y_tmp + 310 + 4*110
		ye := y_tmp + 310 + 1*110
		robotgo.MoveMouse(x, ys)
		robotgo.MouseToggle(`down`, `left`)
		robotgo.Sleep(1)
		robotgo.MoveMouse(x, ye)
		robotgo.Sleep(1)
		robotgo.MouseToggle(`up`, `left`)
		choeseBoss(seq - 3)
		robotgo.Sleep(1)
	}
}

package main

import (
	"github.com/go-vgo/robotgo"
)

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
			haveOneImgsLeft(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"))
		}

		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("gameLogo.png"), getSystemImg("mainMission.png"), getSystemImg("stroke.png"), getSystemImg("itemExchange.png"), getSystemImg("multiplayer.png"), getSystemImg("YES.png"), getSystemImg("OK.png"), getSystemImg("dayGift.png"), getSystemImg("dayClose.png"), getSystemImg("notRecruit.png"), getSystemImg("great.png"), getSystemImg("goPaly.png"), getSystemImg("stop.png")},
			func(x, y int) {
				leftMouseClick(x, y)
				haveOneImgsLeft(10, 0.05, true, getSystemImg("joinMain.png"))
			},
			func(x, y int) {
				if checkpointConfig.Type == "boss" {
					haveOneImgsLeft(10, 0.05, true, getSystemImg("joinBoss.png"))
				} else if checkpointConfig.Type == "activity" {
					haveOneImgsLeft(10, 0.05, true, getSystemImg("joinActivity.png"))
				}
			},
			func(x, y int) { choeseBoss(checkpointConfig.Number) },
			func(x, y int) { choeseDifficulty(checkpointConfig.Difficulty) },
			func(x, y int) { leftMouseClick(x, y) },
			func(x, y int) { leftMouseClick(x, y) },
			func(x, y int) { leftMouseClick(x, y) },
			func(x, y int) { leftMouseClick(x, y) },
			func(x, y int) { leftMouseClick(x, y) },
			func(x, y int) {
				haveOneImgsLeft(5, 0.01, false, getSystemImg("startRaising.png"))
				haveOneImgsLeft(1, 0.01, false, getSystemImg("YrandomRecruitment.png"))
				haveOneImgsLeft(1, 0.01, false, getSystemImg("YfollowersPublic.png"))
				haveOneImgsLeft(5, 0.01, false, getSystemImg("goRecruit.png"))
			},
			func(x, y int) {
				haveOneImgsLeft(5, 0.01, false, getSystemImg("startRaising.png"))
				haveOneImgsLeft(5, 0.01, false, getSystemImg("NrandomRecruitment.png"))
				haveOneImgsLeft(5, 0.01, false, getSystemImg("NfollowersPublic.png"))
				haveOneImgsLeft(5, 0.01, false, getSystemImg("goRecruit.png"))
			},
			func(x, y int) {
				haveOneImgsLeft(5, 0.01, false, getSystemImg("goGame.png"))
			},
			func(x, y int) {
				leftMouseClick(x, y)
				haveOneImgsLeft(5, 0.01, false, getSystemImg("exit.png"))
				haveOneImgsLeft(5, 0.01, false, getSystemImg("exitYes.png"))
			})

		// robotgo.Sleep(3)
	}
}

func choeseBoss(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + 310 + seq*110
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
	y := y_tmp + 200 + seq*110
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

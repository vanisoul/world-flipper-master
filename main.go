package main

import (
	"github.com/go-vgo/robotgo"
)

var yDifficulty = 0
var yBoss = 0
var imgBoss = "remaining.pngOrstroke.png"
var imgDifficulty = "updateList.pngOritemExchange.png"

func main() {

	adbinit(1)

	//初始化ID
	infoID()
	tmpAuto := ""            //紀錄auto狀態 如果是auto模式才有用到
	choeseBossSeq := 0       //這次選擇的關卡
	choeseDifficultySeq := 0 //這次選擇的關卡
	choseAuto := false       //紀錄這次有沒有被改變狀態
	tmpcheckpointConfig, _ := LoadCheckpointConfig()
	tmpcheckpointConfig.Type = "aaaaaaaaaaaaa" //故意改變讓一開始進入回主選單
	tmpRFrequency := 0
	for {
		//如果config有變動 需要重新回到主頁
		checkpointConfig, _ := LoadCheckpointConfig()
		if tmpcheckpointConfig != checkpointConfig || choseAuto {
			tmpcheckpointConfig = checkpointConfig
			if tmpAuto != "repalay" {
				tmpAuto = checkpointConfig.Type
			}
			haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("startRaising.png"), getSystemImg("goGame3.png"), getSystemImg("goGame4.png")},
				func(x, y int) {
					haveOneImgsClick(5, 0.05, true, getSystemImg("return.png"))
					haveOneImgsClick(5, 0.05, true, getSystemImg("disband.png"))
				},
				func(x, y int) {
					haveOneImgsClick(5, 0.05, true, getSystemImg("return.png"))
				},
				func(x, y int) {
					haveOneImgsClick(5, 0.05, true, getSystemImg("return.png"))
				})
			haveOneImgsClick(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"))
			robotgo.Sleep(8)
			choseAuto = false
		}

		//開啟遊戲
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("gameLogo.png"), getSystemImg("joinMain.png"), getSystemImg("mainMission.png"), getSystemImg(imgBoss), getSystemImg(imgDifficulty), getSystemImg("multiplayer.png"), getSystemImg("YES.png"), getSystemImg("OK.png"), getSystemImg("dayGift.png"), getSystemImg("dayClose.png"), getSystemImg("notRecruit.png"), getSystemImg("great.png"), getSystemImg("goPaly.png"), getSystemImg("stop.png"), getSystemImg("exitHalfway.png"), getSystemImg("errorOK.png"), getSystemImg("SKIP.png"), getSystemImg("dayGift2.png"), getSystemImg("gameStop.png"), getSystemImg("errorClose.png"), getSystemImg("fiveStar.png"), getSystemImg("fourStar.png"), getSystemImg("threeStar.png"), getSystemImg("NEW.png"), getSystemImg("OKupdateVertion.png"), getSystemImg("OKupdateVertion2.png"), getSystemImg("playUpdate.png"), getSystemImg("startGamePlay.png"), getSystemImg("updateVertion.png"), getSystemImg("goGame3.png"), getSystemImg("goGame4.png"), getSystemImg("next1.png"), getSystemImg("next2.png"), getSystemImg("next3.png"), getSystemImg("next4.png"), getSystemImg("rePlay.png"), getSystemImg("gmaeOver.png"), getSystemImg("exitYes.png"), getSystemImg("maintag1.png"), getSystemImg("maintag2.png")},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				now_type := ""
				if checkpointConfig.RFeatures {
					now_type = tmpAuto
				} else {
					now_type = checkpointConfig.Type
				}
				if now_type == "boss" {
					haveOneImgsClick(10, 0.05, true, getSystemImg("joinBoss.png"))
					imgBoss = "stroke.png"
					imgDifficulty = "itemExchange.png"
					yDifficulty = 200
					yBoss = 310
					choeseBossSeq = checkpointConfig.FNumber
					choeseDifficultySeq = checkpointConfig.FDifficulty
				} else if now_type == "activity" {
					haveOneImgsClick(10, 0.05, true, getSystemImg("joinActivity.png"))
					imgBoss = "remaining.png"
					imgDifficulty = "updateList.png"
					yDifficulty = 310
					yBoss = 200
					choeseBossSeq = checkpointConfig.FNumber
					choeseDifficultySeq = checkpointConfig.FDifficulty
				} else if now_type == "repalay" {
					haveOneImgsClick(10, 0.05, true, getSystemImg("joinActivity.png"))
					imgBoss = "remaining.png"
					imgDifficulty = "gameProblem.png"
					yDifficulty = 200
					yBoss = 200
					choeseBossSeq = checkpointConfig.RNumber
					choeseDifficultySeq = checkpointConfig.RDifficulty
				}
			},
			func(x, y int) { choeseBoss(choeseBossSeq) },
			func(x, y int) { choeseDifficulty(choeseDifficultySeq) },
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("fullOfEnergy.png")}, func() {
					if checkpointConfig.RFeatures {
						tmpRFrequency = 0
						tmpAuto = "repalay"
						choseAuto = true
					}
				}, func() {
					haveAllImgsExecFunc(1, 0.01, false, []string{getSystemImg("onlyEachOther.png")}, func() {}, func() {
						haveOneImgsClick(5, 0.01, false, getSystemImg("startRaising.png"))
						haveOneImgsClick(1, 0.01, false, getSystemImg("YrandomRecruitment.png"))
						haveOneImgsClick(1, 0.01, false, getSystemImg("YfollowersPublic.png"))
						haveOneImgsClick(5, 0.01, false, getSystemImg("goRecruit.png"))
					})
				})
			},
			func(x, y int) {
				haveAllImgsExecFunc(1, 0.01, false, []string{getSystemImg("recruitAll.png")}, func() {}, func() {
					haveOneImgsClick(2, 0.01, false, getSystemImg("startRaising.png"))
					haveOneImgsClick(1, 0.01, false, getSystemImg("NrandomRecruitment.png"))
					haveOneImgsClick(1, 0.01, false, getSystemImg("NfollowersPublic.png"))
					haveOneImgsClick(2, 0.01, false, getSystemImg("goRecruit.png"))
				})
			},
			func(x, y int) {
				haveOneImgsClick(1, 0.01, false, getSystemImg("goGame.png"))
				haveOneImgsClick(1, 0.01, false, getSystemImg("goGame2.png"))
			},
			func(x, y int) {
				if tmpAuto != "repalay" {
					mouseClick(x, y)
					haveOneImgsClick(5, 0.01, false, getSystemImg("exit.png"))
					haveOneImgsClick(5, 0.01, false, getSystemImg("exitYes.png"))
				}
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				if tmpRFrequency >= checkpointConfig.RFrequency && tmpAuto == "repalay" {
					tmpRFrequency = 0
					tmpAuto = checkpointConfig.Type
					choseAuto = true
				} else {
					mouseClick(x, y)
					tmpRFrequency = tmpRFrequency + 1
					robotgo.Sleep(10)
				}
			},
			func(x, y int) {
				if tmpRFrequency >= checkpointConfig.RFrequency && tmpAuto == "repalay" {
					tmpRFrequency = 0
					tmpAuto = checkpointConfig.Type
					choseAuto = true
				} else {
					mouseClick(x, y)
					tmpRFrequency = tmpRFrequency + 1
					robotgo.Sleep(10)
				}
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				haveOneImgsClick(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"), getSystemImg("main7.png"))
				robotgo.Sleep(8)
			},
			func(x, y int) {
				haveOneImgsClick(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"), getSystemImg("main7.png"))
				robotgo.Sleep(8)
			})

	}
}

func choeseBoss(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yBoss + seq*110
	if seq < 6 {
		mouseClick(x, y)
		robotgo.Sleep(3)
	} else {
		ys := y_tmp + 310 + 4*110
		ye := y_tmp + 310 + 1*110
		AdbShellInputSwipe(x, ys, x, ye)
		choeseBoss(seq - 3)
		robotgo.Sleep(1)
	}
}

func choeseDifficulty(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yDifficulty + seq*110
	if seq < 6 {
		mouseClick(x, y)
		robotgo.Sleep(3)
	} else {
		ys := y_tmp + 310 + 2*110
		ye := y_tmp + 310 + 1*110
		AdbShellInputSwipe(x, ys, x, ye)
		choeseDifficulty(seq - 1)
		robotgo.Sleep(1)
	}
}

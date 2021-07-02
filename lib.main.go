package main

import "github.com/go-vgo/robotgo"

func addGameLogs(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "gameLogo.png", funcs)
	return
}
func addJoinMain(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinMain.png", funcs)
	return
}

func addMainMission(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("mainMission.png"))
	funcs = append(funcs, func(x, y int) {
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
			yBoss = 460
			yEvery = 175
			yDifficulty = 310
			choeseBossSeq = checkpointConfig.FNumber
			choeseDifficultySeq = checkpointConfig.FDifficulty
		} else if now_type == "activity" {
			haveOneImgsClick(10, 0.05, true, getSystemImg("joinActivity.png"))
			imgBoss = "remaining.png"
			imgDifficulty = "updateList.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 460
			choeseBossSeq = checkpointConfig.FNumber
			choeseDifficultySeq = checkpointConfig.FDifficulty
		} else if now_type == "repalay" {
			haveOneImgsClick(10, 0.05, true, getSystemImg("joinActivity.png"))
			imgBoss = "remaining.png"
			imgDifficulty = "gameProblem.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 310
			choeseBossSeq = checkpointConfig.RNumber
			choeseDifficultySeq = checkpointConfig.RDifficulty
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addImgBoss(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg(imgBoss))
	funcs = append(funcs, func(x, y int) {
		choeseBoss(choeseBossSeq)
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addFullOfEnergy(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("fullOfEnergy.png"))
	funcs = append(funcs, func(x, y int) {
		status = 0
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addJoinActivity(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinActivity.png", funcs)
	return
}

func addJoinBoss(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinBoss.png", funcs)
	return
}

func addGoGameOnly(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("goGame3.png"))
	funcs = append(funcs, func(x, y int) {
		if tmpRFrequency > checkpointConfig.RFrequency && checkpointConfig.Type != "repalay" {
			status = 0
		} else {
			tmpRFrequency++
			mouseClick(x, y)
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addGoGameMaze(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("goGameMaze.png"))
	funcs = append(funcs, func(x, y int) {
		if tmpRFrequency > checkpointConfig.RFrequency && checkpointConfig.Type != "repalay" {
			status = 0
		} else {
			tmpRFrequency++
			mouseClick(x, y)
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addImgDifficulty(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg(imgDifficulty))
	funcs = append(funcs, func(x, y int) {
		choeseDifficulty(choeseDifficultySeq)
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addMultiplayer(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("multiplayer.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addYES(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "YES.png", funcs)
	return
}

func addOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "OK.png", funcs)
	return
}

func addGameOverOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "GameOverOK.png", funcs)
	return
}

func addDayGift(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "dayGift.png", funcs)
	return
}

func addDayClose(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "dayClose.png", funcs)
	return
}

func addNotRecruit(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("notRecruit.png"))
	funcs = append(funcs, func(x, y int) {
		haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("onlyEachOther.png")}, func() {
			status = 66
		}, func() {
			haveOneImgsClick(5, 0.01, false, getSystemImg("startRaising.png"))
			haveOneImgsClick(1, 0.01, false, getSystemImg("YrandomRecruitment.png"))
			haveOneImgsClick(1, 0.01, false, getSystemImg("YfollowersPublic.png"))
			haveOneImgsClick(5, 0.01, false, getSystemImg("goRecruit.png"))
		})
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addOnlyEachOther(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("onlyEachOther.png"))
	funcs = append(funcs, func(x, y int) {
		status = 66
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addNotFullOfEnergyMain(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("mainMission.png"))
	funcs = append(funcs, func(x, y int) {
		if checkpointConfig.Type == "boss" {
			imgBoss = "stroke.png"
			imgDifficulty = "itemExchange.png"
			yBoss = 460
			yEvery = 175
			yDifficulty = 310
			choeseBossSeq = checkpointConfig.FNumber
			choeseDifficultySeq = checkpointConfig.FDifficulty
			status = 6
		} else if checkpointConfig.Type == "activity" {
			imgBoss = "remaining.png"
			imgDifficulty = "updateList.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 460
			choeseBossSeq = checkpointConfig.FNumber
			choeseDifficultySeq = checkpointConfig.FDifficulty
			status = 6
		} else if checkpointConfig.Type == "repalay" {
			imgBoss = "remaining.png"
			imgDifficulty = "gameProblem.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 310
			choeseBossSeq = checkpointConfig.RNumber
			choeseDifficultySeq = checkpointConfig.RDifficulty
			status = 5
		} else {
			imgBoss = "isNotFound"
			imgDifficulty = "isNotFound"
			yBoss = -1
			yEvery = -1
			yDifficulty = -1
			choeseBossSeq = -1
			status = 0
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addFullOfEnergyMain(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("fullOfEnergy.png"))
	funcs = append(funcs, func(x, y int) {
		if checkpointConfig.RFeatures {
			imgBoss = "remaining.png"
			imgDifficulty = "gameProblem.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 310
			choeseBossSeq = checkpointConfig.RNumber
			choeseDifficultySeq = checkpointConfig.RDifficulty
			tmpRFrequency = 0
			status = 5
		} else if checkpointConfig.Type == "boss" {
			imgBoss = "stroke.png"
			imgDifficulty = "itemExchange.png"
			yBoss = 460
			yEvery = 175
			yDifficulty = 310
			choeseBossSeq = checkpointConfig.FNumber
			choeseDifficultySeq = checkpointConfig.FDifficulty
			status = 7
		} else if checkpointConfig.Type == "activity" {
			imgBoss = "remaining.png"
			imgDifficulty = "updateList.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 460
			choeseBossSeq = checkpointConfig.FNumber
			choeseDifficultySeq = checkpointConfig.FDifficulty
			status = 6
		} else if checkpointConfig.Type == "repalay" {
			imgBoss = "remaining.png"
			imgDifficulty = "gameProblem.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 310
			choeseBossSeq = checkpointConfig.RNumber
			choeseDifficultySeq = checkpointConfig.RDifficulty
			status = 5
		} else {
			imgBoss = "isNotFound"
			imgDifficulty = "isNotFound"
			yBoss = -1
			yEvery = -1
			yDifficulty = -1
			choeseBossSeq = -1
			status = 0
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addGreat(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("great.png"))
	funcs = append(funcs, func(x, y int) {
		haveAllImgsExecFunc(1, 0.01, false, []string{getSystemImg("recruitAll.png")}, func() {}, func() {
			haveOneImgsClick(2, 0.01, false, getSystemImg("startRaising.png"))
			haveOneImgsClick(1, 0.01, false, getSystemImg("NrandomRecruitment.png"))
			haveOneImgsClick(1, 0.01, false, getSystemImg("NfollowersPublic.png"))
			haveOneImgsClick(2, 0.01, false, getSystemImg("goRecruit.png"))
		})
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addGoPaly(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("goPaly.png"))
	funcs = append(funcs, func(x, y int) {
		haveOneImgsClick(1, 0.01, false, getSystemImg("goGame.png"))
	})
	resStrs = strs
	resFuncs = funcs
	return
}
func addStop(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("stop.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
		haveOneImgsClick(5, 0.01, false, getSystemImg("exit.png"))
		haveOneImgsClick(5, 0.01, false, getSystemImg("exitYes.png"))
		status = 6

	})
	resStrs = strs
	resFuncs = funcs
	return
}
func addGoGame3(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("goGame3.png"))
	funcs = append(funcs, func(x, y int) {
		if tmpRFrequency >= checkpointConfig.RFrequency && tmpAuto == "repalay" {
			tmpRFrequency = 0
			tmpAuto = checkpointConfig.Type
			choseAuto = true
		} else {
			mouseClick(x, y)
			tmpRFrequency = tmpRFrequency + 1
			robotgo.Sleep(10)
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}
func addNext1(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "next1.png", funcs)
	return
}
func addRePlay(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "rePlay.png", funcs)
	return
}
func addExitYes(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("exitYes.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
		status = 6
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addExitOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("OK.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
		status = 6
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addNetworkerrorOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "networkerrorOK.png", funcs)
	return
}

func addStartRaising(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("startRaising.png"))
	funcs = append(funcs, func(x, y int) {
		haveOneImgsClick(5, 0.05, true, getSystemImg("return.png"))
		haveOneImgsClick(5, 0.05, true, getSystemImg("disband.png"))
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addRgoGame3(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("goGame3.png"))
	funcs = append(funcs, func(x, y int) {
		haveOneImgsClick(5, 0.05, true, getSystemImg("return.png"))
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addMain2(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "main2.png", funcs)
	return
}
func addGmaeOver(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "gmaeOver.png", funcs)
	return
}
func addExitHalfway(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "exitHalfway.png", funcs)
	return
}

func addMainMissionMainOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("mainMission.png"))
	funcs = append(funcs, func(x, y int) {
		status = 2
	})
	resStrs = strs
	resFuncs = funcs
	return
}

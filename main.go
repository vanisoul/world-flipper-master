package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		//初始化設定 未來要判斷視窗位置
		infoScreen()
		//初始化ID
		infoID()
		//首抽流程 其中一步失敗都會出來
		firstCard()
	}
}

func firstCard() {
	//清除資料 確保全新
	mainDataSucc := checkImgClickOtherImg(GetSystemImg("main.png"), GetSystemImg("list.png"), nil, 0.1)
	if !mainDataSucc {
		reStart()
		return
	}
	delDataSucc := leftMouseClickImg(GetSystemImg("delData.png"))
	if !delDataSucc {
		reStart()
		return
	}
	redOK1Succ := leftMouseClickImg(GetSystemImg("redOK.png"))
	if !redOK1Succ {
		reStart()
		return
	}
	robotgo.Sleep(2)
	redOK2Succ := leftMouseClickImg(GetSystemImg("redOK2.png"))
	if !redOK2Succ {
		reStart()
		return
	}
	robotgo.Sleep(10)

	//進入遊戲
	mainSucc := leftMouseClickImg(GetSystemImg("main.png"), 0.1)
	if !mainSucc {
		reStart()
		return
	}

	//同意規章
	agreeSucc := checkImgClickOtherImg(GetSystemImg("agree.png"), GetSystemImg("agreeOK.png"), nil, 0.1)
	if !agreeSucc {
		reStart()
		return
	}

	//是否遊玩過遊戲
	playGameSucc := checkImgClickOtherImg(GetSystemImg("playGame.png"), GetSystemImg("playGameYes.png"), nil, 0.1)
	if !playGameSucc {
		reStart()
		return
	}

	//跳過新手教學
	skipGameSucc := checkImgClickOtherImg(GetSystemImg("skipGame.png"), GetSystemImg("skipOK.png"), nil, 0.1)
	if !skipGameSucc {
		reStart()
		return
	}

	//故事大綱
	OK2Succ := leftMouseClickImg(GetSystemImg("OK2.png"))
	if !OK2Succ {
		reStart()
		return
	}

	//教學畫面 一直下一步
	next1 := leftMouseClickImg(GetSystemImg("next.png"))
	if !next1 {
		reStart()
		return
	}
	robotgo.Sleep(1)
	next2 := leftMouseClickImg(GetSystemImg("next.png"))
	if !next2 {
		reStart()
		return
	}
	robotgo.Sleep(1)
	next3 := leftMouseClickImg(GetSystemImg("next.png"))
	if !next3 {
		reStart()
		return
	}
	robotgo.Sleep(1)
	next4 := leftMouseClickImg(GetSystemImg("next.png"))
	if !next4 {
		reStart()
		return
	}
	next5 := leftMouseClickImg(GetSystemImg("closenext.png"))
	if !next5 {
		reStart()
		return
	}

	//輸入名稱
	enterNameSucc := checkImgClickOtherImg(GetSystemImg("enterName.png"), GetSystemImg("OK.png"))
	if !enterNameSucc {
		reStart()
		return
	}
	checkNameSucc := checkImgClickOtherImg(GetSystemImg("checkName.png"), GetSystemImg("checkName2.png"))
	if !checkNameSucc {
		reStart()
		return
	}

	//點選轉蛋
	capsuleTipSucc := checkImgClickOtherImg(GetSystemImg("capsuleTip.png"), GetSystemImg("capsule.png"))
	if !capsuleTipSucc {
		reStart()
		return
	}

	//試著抽一抽
	tryCapsuleTipSucc := checkImgClickOtherImg(GetSystemImg("tryCapsuleTip.png"), GetSystemImg("oneCapsule.png"))
	if !tryCapsuleTipSucc {
		reStart()
		return
	}
	checkCapsuleSucc := leftMouseClickImg(GetSystemImg("checkCapsule.png"))
	if !checkCapsuleSucc {
		reStart()
		return
	}

	// 截圖下來 第一隻肯定都有NEW 沒有代表出錯
	// 並且如果有NEW 就解析人物 大圖
	tenCap()

	//回到轉蛋首頁
	goBackCapsuleSucc := leftMouseClickImg(GetSystemImg("goBackCapsule.png"))
	if !goBackCapsuleSucc {
		reStart()
		return
	}

	//上市禮物1500寶石
	giftSucc := checkImgClickOtherImg(GetSystemImg("gift.png"), GetSystemImg("closeGift.png"))
	if !giftSucc {
		reStart()
		return
	}
	//又跳人物
	tenCap()

	//切換到首頁
	priSucc := leftMouseClickImgMany(0.2, []string{GetSystemImg("primary1.png"), GetSystemImg("primary2.png"), GetSystemImg("primary3.png"), GetSystemImg("primary4.png")}...)
	if !priSucc {
		reStart()
		return
	}

	//獎勵領取
	checkGiftSucc := leftMouseClickImg(GetSystemImg("checkGift.png"))
	if !checkGiftSucc {
		reStart()
		return
	}

	//天獎勵領取
	dayCheckGiftSucc := leftMouseClickImg(GetSystemImg("dayCheckGift.png"))
	if !dayCheckGiftSucc {
		reStart()
		return
	}

	//點信箱
	mailSucc := leftMouseClickImg(GetSystemImg("mail.png"))
	if !mailSucc {
		reStart()
		return
	}

	//全部領取
	allTakeSucc := leftMouseClickImg(GetSystemImg("allTake.png"))
	if !allTakeSucc {
		reStart()
		return
	}
	takeYesSucc := leftMouseClickImg(GetSystemImg("takeYes.png"))
	if !takeYesSucc {
		reStart()
		return
	}

	//關閉信箱
	okmailSucc := checkImgClickOtherImg(GetSystemImg("okmail.png"), GetSystemImg("closemail.png"))
	if !okmailSucc {
		reStart()
		return
	}

	// 點選轉蛋
	capsuleSucc := leftMouseClickImgMany(0.01, GetSystemImg("capsule2.png"), GetSystemImg("capsule.png"))
	if !capsuleSucc {
		reStart()
		return
	}

	//十抽
	tenCapsuleSucc := leftMouseClickImg(GetSystemImg("tenCapsule.png"))
	if !tenCapsuleSucc {
		reStart()
		return
	}
	checkCapSucc := leftMouseClickImg(GetSystemImg("checkCapsule.png"))
	if !checkCapSucc {
		reStart()
		return
	}
	//處理那10抽
	tenCap()

	//點選轉蛋
	capsuleSucc1 := leftMouseClickImgMany(0.01, GetSystemImg("capsule2.png"), GetSystemImg("capsule.png"))
	if !capsuleSucc1 {
		reStart()
		return
	}

	//十抽
	tenCapsuleSucc1 := leftMouseClickImg(GetSystemImg("tenCapsule.png"))
	if !tenCapsuleSucc1 {
		reStart()
		return
	}
	checkCapSucc1 := leftMouseClickImg(GetSystemImg("checkCapsule.png"))
	if !checkCapSucc1 {
		reStart()
		return
	}
	//處理那10抽
	tenCap()

	//點選轉蛋
	capsuleSucc1 = leftMouseClickImgMany(0.01, GetSystemImg("capsule2.png"), GetSystemImg("capsule.png"))
	if !capsuleSucc1 {
		reStart()
		return
	}

	//十抽
	tenCapsuleSucc2 := leftMouseClickImg(GetSystemImg("tenCapsule.png"))
	if !tenCapsuleSucc2 {
		reStart()
		return
	}
	checkCapSucc2 := leftMouseClickImg(GetSystemImg("checkCapsule.png"))
	if !checkCapSucc2 {
		reStart()
		return
	}
	//處理那10抽
	tenCap()

	//選單
	menuSucc := leftMouseClickImg(GetSystemImg("menu.png"))
	if !menuSucc {
		reStart()
		return
	}

	//繼承 inherit
	inheritSucc := leftMouseClickImg(GetSystemImg("inherit.png"))
	if !inheritSucc {
		reStart()
		return
	}
	//繼承資料 inheritText
	inheritTextSucc := leftMouseClickImg(GetSystemImg("inheritText.png"))
	if !inheritTextSucc {
		reStart()
		return
	}

	//繼承資料 inheritNext
	inheritNextSucc := leftMouseClickImg(GetSystemImg("inheritNext.png"))
	if !inheritNextSucc {
		reStart()
		return
	}

	//idandpwd
	idandpwdSucc := leftMouseClickImg(GetSystemImg("idandpwd.png"))
	if !idandpwdSucc {
		reStart()
		return
	}
	//idandpwdOK
	idandpwdOKSucc := leftMouseClickImg(GetSystemImg("idandpwdOK.png"))
	if !idandpwdOKSucc {
		reStart()
		return
	}
	//inheritNextText
	inheritNextTextSucc := leftMouseClickImg(GetSystemImg("inheritNextText.png"))
	if !inheritNextTextSucc {
		reStart()
		return
	}
	//enterPWD
	//設定密碼
	enterPWDSucc, ep1_x, ep1_y := whilescreen(GetSystemImg("enterPWD.png"), 0.1)
	if !enterPWDSucc {
		reStart()
		return
	}
	pwd = fmt.Sprintf("pWd%s", thisID)
	textLocation(ep1_x, ep1_y, pwd)
	enterPWD2Succ, ep2_x, ep2_y := whilescreen(GetSystemImg("enterPWD2.png"), 0.1)
	if !enterPWD2Succ {
		reStart()
		return
	}
	textLocation(ep2_x, ep2_y, pwd)

	//隱私權OK
	inheritOKSucc := leftMouseClickImg(GetSystemImg("inheritOK.png"))
	if !inheritOKSucc {
		reStart()
		return
	}

	//inheritOKnext
	inheritOKnextSucc := leftMouseClickImg(GetSystemImg("inheritOKnext.png"))
	if !inheritOKnextSucc {
		reStart()
		return
	}
	robotgo.Sleep(5)
	savescreen(thisID)

	//endOK
	endOKSucc := leftMouseClickImg(GetSystemImg("endOK.png"))
	if !endOKSucc {
		reStart()
		return
	}

	outputRoleLog()
	reStart()
}

func tenCap() {
	robotgo.Sleep(10)
	for {
		bigRoleMain()
		//選單
		menuSucc, _, _ := whilescreenbase(GetSystemImg("menu.png"), 1, 0.01)
		if menuSucc {
			break
		}
		repOKSucc, rx, ry := whilescreenbase(GetSystemImg("repOK.png"), 1, 0.01)
		if repOKSucc {
			robotgo.MoveMouse(rx, ry)
			leftMouseClick()
			break
		}
		okNewTipOKSucc, rx, ry := whilescreenbase(GetSystemImg("okNewTipOK.png"), 1, 0.01)
		if okNewTipOKSucc {
			robotgo.MoveMouse(rx, ry)
			leftMouseClick()
			break
		}

	}
}

func bigRoleMain() bool {
	newSucc := checkImgSaveImg(GetSystemImg("NEW.png"), 1, 0.1)
	if !newSucc {
		leftMouseClick()
		return false
	} else {
		analyzeRoleImg("big")
		_, x, y := whilescreenbase(GetSystemImg("NEW.png"), 1, 0.1)
		robotgo.MoveMouse(x, y)
		leftMouseClick()
		return true
	}
}

//遇到不如預期時 離開這次流程前 需要做的事情
func reStart() {
	//關閉遊戲
	leftMouseClickImg(GetSystemImg("appSmall.png"), 0.1)
	_, x, y := whilescreen(GetSystemImg("smallappLogo.png"), 0.05)
	robotgo.MoveMouse(x, y)
	robotgo.MouseToggle(`down`, `left`)
	robotgo.Sleep(1)
	robotgo.MoveMouseSmooth(x+1500, y)
	robotgo.Sleep(1)
	robotgo.MouseToggle(`up`, `left`)

	//打開遊戲
	leftMouseClickImg(GetSystemImg("appLogo.png"))
	robotgo.Sleep(15)
}

// 確認初始化視窗位置是否正確
func checkinfoImg() {
	//初始化設定 未來要判斷視窗位置
	infoScreen()
	savescreen()
}

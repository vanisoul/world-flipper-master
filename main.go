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

	// 第一抽
	capSucc1 := caping()
	if !capSucc1 {
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
	iolesaSucc, iex, iey := whilescreenbase(GetSystemImg("iolesa.png"), 25, 0.01)
	if iolesaSucc {
		robotgo.MoveMouse(iex, iey)
		leftMouseClick()
	}
	repOKSucc, rx, ry := whilescreenbase(GetSystemImg("iolesaOK.png"), 25, 0.01)
	if repOKSucc {
		robotgo.MoveMouse(rx, ry)
		leftMouseClick()
	}

	//切換到首頁
	priSucc := leftMouseClickImgMany(0.4, []string{GetSystemImg("primary1.png"), GetSystemImg("primary2.png"), GetSystemImg("primary3.png"), GetSystemImg("primary4.png")}...)
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
	dayCheckGiftSucc := leftMouseClickImg(GetSystemImg("dayCheckGift.png"), 0.1)
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

	//幾次十抽
	tcSucc1 := tenCap()
	if !tcSucc1 {
		reStart()
		return
	}

	tcSucc2 := tenCap()
	if !tcSucc2 {
		reStart()
		return
	}

	//選單
	menuSucc := leftMouseClickImg(GetSystemImg("menu.png"))
	if !menuSucc {
		reStart()
		return
	}

	//繼承 inherit
	inheritSucc := leftMouseClickImg(GetSystemImg("inherit.png"))
	if !inheritSucc {
		menuSucc := leftMouseClickImg(GetSystemImg("menu.png"))
		if !menuSucc {
			reStart()
			return
		}
		robotgo.Sleep(1)
		leftMouseClickImg(GetSystemImg("inherit.png"))
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

func tenCap() bool {

	//點選轉蛋
	leftMouseClickImgMany(0.1, GetSystemImg("capsule2.png"), GetSystemImg("capsule.png"))

	robotgo.Sleep(3)
	//選池
	tgCapSucc, tgx, tgy := whilescreen(GetSystemImg("tgCap.png"), 0.1)
	if tgCapSucc {
		robotgo.MoveMouse(tgx, tgy)
		leftMouseClick()
		robotgo.Sleep(5)
	}

	//十抽
	tenCapsuleSucc := leftMouseClickImg(GetSystemImg("tenCapsule.png"), 0.2)
	if !tenCapsuleSucc {
		return false
	}
	checkCapSucc := leftMouseClickImg(GetSystemImg("checkCapsule.png"), 0.1)
	if !checkCapSucc {
		return false
	}
	//處理那10抽
	capSucc2 := caping()
	return capSucc2
}

func caping() bool {
	count := 300
	for count > 0 {
		leftMouseClick()
		leftMouseClick()
		leftMouseClick()
		leftMouseClick()
		recapSucc, _, _ := whilescreenbase(GetSystemImg("recap.png"), 1, 0.05)
		if recapSucc {
			robotgo.Sleep(1)
			recapSucc2, rx, ry := whilescreenbase(GetSystemImg("recap.png"), 1, 0.05)
			if recapSucc2 {
				robotgo.MoveMouse(rx, ry)
				leftMouseClick()
			}
		}
		recapSucc1, _, _ := whilescreenbase(GetSystemImg("yysa.png"), 1, 0.1)
		if recapSucc1 {
			robotgo.Sleep(1)
			recapSucc2, rx, ry := whilescreenbase(GetSystemImg("yysa.png"), 1, 0.1)
			if recapSucc2 {
				robotgo.MoveMouse(rx, ry)
				leftMouseClick()
			}
		}
		//回到轉蛋首頁
		bakCapSucc, bcx, bcy := whilescreenbase(GetSystemImg("goBackCapsule.png"), 1, 0.1)
		if bakCapSucc {
			robotgo.Sleep(1)
			recapSucc, rx, ry := whilescreenbase(GetSystemImg("recap.png"), 1, 0.1)
			if recapSucc {
				robotgo.MoveMouse(rx, ry)
				leftMouseClick()
			}
			robotgo.Sleep(1)
			savescreen(thisID)
			robotgo.MoveMouse(bcx, bcy)
			leftMouseClick()
			return true
		}
		robotgo.Sleep(1)
		count = count - 1
	}
	return false
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
	savescreen(thisID)

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

package main

func main() {

	//初始化設定 未來要判斷視窗位置
	infoScreen()
	//初始化ID
	infoID()
	//首抽流程 其中一步失敗都會出來
	firstCard()
}

func firstCard() {
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

	//教學畫面 一直下一步
	next1 := leftMouseClickImg(GetSystemImg("next.png"))
	if !next1 {
		reStart()
		return
	}
	next2 := leftMouseClickImg(GetSystemImg("next.png"))
	if !next2 {
		reStart()
		return
	}
	next3 := leftMouseClickImg(GetSystemImg("next.png"))
	if !next3 {
		reStart()
		return
	}
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
	checkNameSucc := checkImgClickOtherImg(GetSystemImg("checkName.png"), GetSystemImg("OK.png"))
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
	newSucc := checkImgSaveImg(GetSystemImg("NEW.png"))
	if !newSucc {
		reStart()
		return
	} else {
		analyzeRoleImg("big")
		leftMouseClickImg(GetSystemImg("NEW.png"))
	}

	okNewTipSucc := checkImgClickOtherImg(GetSystemImg("okNewTip.png"), GetSystemImg("OK.png"))
	if !okNewTipSucc {
		reStart()
		return
	}

	//回到轉蛋首頁
	goBackCapsuleSucc := leftMouseClickImg(GetSystemImg("goBackCapsule.png"))
	if !goBackCapsuleSucc {
		reStart()
		return
	}

}

//遇到不如預期時 離開這次流程前 需要做的事情
func reStart() {

}

// 確認初始化視窗位置是否正確
func checkinfoImg() {
	//初始化設定 未來要判斷視窗位置
	infoScreen()
	savescreen()
}

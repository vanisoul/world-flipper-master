package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

var thisID string = "0"

var fiveRole []string = []string{}
var fourRole []string = []string{}
var threeRole []string = []string{}

//初始化這次ID
func infoID() {
	thisID = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	setLog("infoID", "初始化本次刷首抽ID", thisID)
}

//解析圖片 typeImg {big, small} |
// img 圖片 _Ctype_MMBitmapRef 型別
func analyzeRoleImg(typeImg string) {

	if typeImg == "big" {
		roleType, roleNames := analyzeRoleBig()
		setRole(roleType, roleNames...)
	}

	if typeImg == "small" {
		roleType, roleNames := analyzeRoleSmall()
		setRole(roleType, roleNames...)
	}
}

//TODO 大圖小圖解析還沒處理

//給予大圖片 回傳人物名稱 星數
func analyzeRoleBig() (roleType int, roleNames []string) {
	bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
	defer robotgo.FreeBitmap(bitmap)

	return
}

//給予多張小圖片 回傳人物名稱 星數
func analyzeRoleSmall() (roleType int, roleNames []string) {
	bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
	defer robotgo.FreeBitmap(bitmap)
	return
}

//給予星數 + roleNames 自動放入該星數內容中
func setRole(roleType int, roleNames ...string) {
	if roleType == 5 {
		setFive(roleNames...)
	}
	if roleType == 4 {
		setFour(roleNames...)
	}
	if roleType == 3 {
		setThree(roleNames...)
	}
}

//紀錄本次五星角色
func setFive(roleNames ...string) {
	for _, roleName := range roleNames {
		fiveRole = append(fiveRole, roleName)
	}
}

//紀錄本次四星角色
func setFour(roleNames ...string) {
	for _, roleName := range roleNames {
		fourRole = append(fourRole, roleName)
	}
}

//紀錄本次三星角色
func setThree(roleNames ...string) {
	for _, roleName := range roleNames {
		threeRole = append(threeRole, roleName)
	}
}

//輸出結算Log
func outputLog() {
	setLog("ID", "本次ID", thisID)
	setLog("fiveRole", "五星角", strings.Join(fiveRole, ", "))
	setLog("fourRole", "四星角", strings.Join(fourRole, ", "))
	setLog("threeRole", "三星角", strings.Join(threeRole, ", "))
}

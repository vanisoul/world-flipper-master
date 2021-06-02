package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
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

//解析圖片 typeImg {big, small} 並自動存入全域變數
func analyzeRoleImg(typeImg string) {

	if typeImg == "big" {
		roleArr := analyzeRoleBig()
		setRole(roleArr)
	}

	if typeImg == "small" {
		roleArr := analyzeRoleSmall()
		setRole(roleArr)
	}
}

//TODO 大圖小圖解析還沒處理

//給予大圖片 回傳人物名稱 星數
func analyzeRoleBig() (res []Role) {
	bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
	defer robotgo.FreeBitmap(bitmap)
	bigRoleConfig, err := LoadBigRoleConfig()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	for _, RoleImg := range bigRoleConfig.RoleImgPair {
		isRole, _, _ := whilescreenbase(getImgBig(RoleImg.ImgName), 2, 0.01)
		if isRole {
			p := Role{RoleType: RoleImg.RoleType, RoleName: RoleImg.RoleName}
			res = append(res, p)
			return
		}
	}

	return
}

//給予多張小圖片 回傳人物名稱 星數
func analyzeRoleSmall() (res []Role) {
	bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
	defer robotgo.FreeBitmap(bitmap)
	smallRoleConfig, err := LoadSmallRoleConfig()
	if err != nil {
		log.Errorf("cannot load config:", err)
		return
	}
	for _, RoleImg := range smallRoleConfig.RoleImgPair {
		isRole, _, _ := whilescreenbase(getImgSmall(RoleImg.ImgName), 2, 0.01)
		if isRole {
			p := Role{RoleType: RoleImg.RoleType, RoleName: RoleImg.RoleName}
			res = append(res, p)
		}
	}
	return
}

//給予星數 + roleNames 自動放入該星數內容中
func setRole(res []Role) {
	for _, v := range res {
		roleType := v.RoleType
		roleName := v.RoleName
		if roleType == 5 {
			setFive(roleName)
		}
		if roleType == 4 {
			setFour(roleName)
		}
		if roleType == 3 {
			setThree(roleName)
		}
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
func outputRoleLog() {
	setLog("ID", "本次ID", thisID)
	setLog("fiveRole", "五星角", strings.Join(fiveRole, ", "))
	setLog("fourRole", "四星角", strings.Join(fourRole, ", "))
	setLog("threeRole", "三星角", strings.Join(threeRole, ", "))
}

type Role struct {
	RoleType int    `mapstructure:"roleType"`
	RoleName string `mapstructure:"roleName"`
}

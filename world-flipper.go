package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
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

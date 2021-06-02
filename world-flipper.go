package main

import (
	"math/rand"
	"strconv"
	"time"
)

var thisID string = "0"

//初始化這次ID
func infoID() {
	thisID = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	setLog("infoID", "初始化本次刷首抽ID", thisID)
}

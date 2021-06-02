package main

import (
	"math/rand"
	"strconv"
	"time"
)

var thisID string = "0"

//看到指定圖片 就截圖存放
func checkImgSaveImg(fullImg string) bool {
	succ := findscreen(fullImg, 0.1)
	if succ {
		path := []string{thisID}
		savescreen(path...)
		return true
	}
	return false
}

//初始化這次ID
func infoID() {
	thisID = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}

package main

var infox int = 0
var infoy int = 0
var infow int = 2000
var infoh int = 2000

//手機板 900*1600

//初始化本程式未來要找的座標範圍 (不找全圖)
func infoScreen() bool {
	succ, x, y := whilescreen(GetSystemImg("noxLogo.png"), 0.05)
	if succ {
		infox = x - 40
		infoy = y - 40
		infow = 600
		infoh = 1050
		return true
	}
	return false
}

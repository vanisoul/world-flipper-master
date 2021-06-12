package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

var rigorousSec = 1

//搜尋一張圖片 符合回傳true,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 秒數可由rigorousSec更改,
//imgFullPaths : 傳入圖片相對全路徑 (多張),
//回傳 : 找到與否 - 找到的是第幾張 - 找到的中心點x - 找到的中心點y
func findOneImages(frequency int, matchNumber float32, rigorous bool, imgFullPaths ...string) (succ bool, index int, x int, y int) {
	log.Infof("findOneImages, frequency: %d, matchNumber: %e, rigorous: %s", frequency, matchNumber, rigorous)
	succ = false
	index = -1
	x = -1
	y = -1

	robotgo.MoveMouse(0, 0)
	new_bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
	defer robotgo.FreeBitmap(new_bitmap)

	rigorous_bitmap := robotgo.CaptureScreen(0, 0, 0, 0)
	if rigorous {
		robotgo.Sleep(rigorousSec)
		rigorous_bitmap = robotgo.CaptureScreen(infox, infoy, infow, infoh)
	}
	defer robotgo.FreeBitmap(rigorous_bitmap)

	for idx, imgFullPath := range imgFullPaths {
		fileExist, imgw, imgh := FileExist(imgFullPath)
		if !fileExist {
			continue
		}
		dst_map := robotgo.OpenBitmap(imgFullPath)
		defer robotgo.FreeBitmap(dst_map)

		fx, fy := robotgo.FindBitmap(dst_map, new_bitmap, matchNumber)
		//如果沒找到就下一張圖
		if fx == -1 && fy == -1 {
			continue
		}
		//如果需要嚴謹 就在找一次 如果沒有 一樣進入下張圖
		if rigorous {
			fx, fy = robotgo.FindBitmap(dst_map, rigorous_bitmap, matchNumber)
			if fx == -1 && fy == -1 {
				continue
			}
		}
		succ = true
		index = idx
		x = fx + (imgw / 2) + infox
		y = fy + (imgh / 2) + infoy
		log.Infof("findOneImages %s is found", imgFullPath)
		setLog("findOneImages", "尋找圖片成功", fmt.Sprintf("img: %s, x:%d, y:%d", imgFullPath, x, y))
		return
	}
	return
}

//搜尋多張圖片 都符合回傳true,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 : 秒數可由rigorousSec修改,
//imgFullPaths : 傳入圖片相對全路徑 (多張),
//回傳 : 找到與否
func findAllImages(frequency int, matchNumber float32, rigorous bool, imgFullPaths ...string) (succ bool) {
	log.Infof("findAllImages, frequency: %d, matchNumber: %e, rigorous: %s", frequency, matchNumber, rigorous)
	succ = false

	filesExist := FilesExist(imgFullPaths...)
	if !filesExist {
		return
	}

	robotgo.MoveMouse(0, 0)
	new_bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
	defer robotgo.FreeBitmap(new_bitmap)

	rigorous_bitmap := robotgo.CaptureScreen(0, 0, 0, 0)
	if rigorous {
		robotgo.Sleep(rigorousSec)
		rigorous_bitmap = robotgo.CaptureScreen(infox, infoy, infow, infoh)
	}
	defer robotgo.FreeBitmap(rigorous_bitmap)

	for _, imgFullPath := range imgFullPaths {
		dst_map := robotgo.OpenBitmap(imgFullPath)
		defer robotgo.FreeBitmap(dst_map)

		fx, fy := robotgo.FindBitmap(dst_map, new_bitmap, matchNumber)
		//如果沒找到就下一張圖
		if fx == -1 && fy == -1 {
			return
		}
		//如果需要嚴謹 就在找一次 如果沒有 一樣進入下張圖
		if rigorous {
			fx, fy = robotgo.FindBitmap(dst_map, rigorous_bitmap, matchNumber)
			if fx == -1 && fy == -1 {
				return
			}
		}
		log.Infof("findAllImages %s is found", imgFullPath)
		setLog("findAllImages", "尋找圖片成功", fmt.Sprintf("img: %s", imgFullPath))
	}
	succ = true
	return
}

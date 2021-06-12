package main

import (
	"image"
	"io"
	"os"

	"github.com/labstack/gommon/log"
)

// 判斷檔案存在
func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// 判斷圖片存在與否
func imgFileExist(path string) (succ bool, w int, h int) {
	file, _ := os.Open(path)
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		file.Close()
		log.Errorf("image file error , is not found %s", path)
		succ = false
		w = -1
		h = -1
		return
	}
	succ = true
	w = c.Width
	h = c.Height
	file.Close()
	return
}

// 判斷多個圖片存在與否
func imgFilesExist(paths ...string) (succ bool) {
	for _, path := range paths {
		succ, _, _ = imgFileExist(path)
		if succ == false {
			return
		}
	}
	return
}

//複製檔案 給路徑 (來源,目的)
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

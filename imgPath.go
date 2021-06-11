package main

import "path"

//給予圖片名稱 自動組出img/system/${name}
func getSystemImg(name string) string {
	return path.Join("img", "system", name)
}

//給予子目錄名稱 以及 圖片名稱 自動組出img/${pathStr}/${name}
func getImg(pathStr string, name string) string {
	return path.Join("img", pathStr, name)
}

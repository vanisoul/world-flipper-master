package main

import "path"

func GetSystemImg(name string) string {
	return path.Join("img", "system", name)
}

func getImgBig(name string) string {
	return path.Join("img", "role", "big", name)
}

func getImgSmall(name string) string {
	return path.Join("img", "role", "small", name)
}

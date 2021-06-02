package main

import "path"

func GetSystemImg(name string) string {
	return path.Join("img", "system", name)
}

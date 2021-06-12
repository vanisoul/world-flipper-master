package main

type List struct {
	contentList []string
}

func (l List) addExampleList() List {
	tmp_List := []string{getSystemImg("A.png")}
	l.contentList = append(l.contentList, tmp_List...)
	return l
}

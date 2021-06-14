package main

type List struct {
	contentList []string
}

func (l List) addCheckMission() List {
	tmp_List := []string{getSystemImg("mainMission.png")}
	l.contentList = append(l.contentList, tmp_List...)
	return l
}

func (l List) addJoinGame() List {
	tmp_List := []string{getSystemImg("gameLogo.png"), getSystemImg("joinMain.png")}
	l.contentList = append(l.contentList, tmp_List...)
	return l
}

func (l List) addExampleList() List {
	tmp_List := []string{getSystemImg("A.png")}
	l.contentList = append(l.contentList, tmp_List...)
	return l
}

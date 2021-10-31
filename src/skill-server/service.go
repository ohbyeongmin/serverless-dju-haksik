package main

import "fmt"

func GetOneMenuReasponse(menu []string, lunchOrDinner LunOrDin) *SkillResponseType {
	var title string
	if lunchOrDinner == LUNCH {
		title = "점심"
	} else {
		title = "저녁"
	}

	itemList := []ICItem{}
	for i, v := range menu {
		itemList = append(itemList, ICItem{Title: fmt.Sprintf("%d", i+1), Description: v})
	}
	itemCard := ItemCardType{
		Head: ICHead{
			Title: title,
		},
		ItemList: itemList,
	}
	outputs := []OutputsType{
		{itemCard},
	}
	template := TemplateType{
		outputs,
	}
	res := SkillResponseType{
		Version:  "2.0",
		Template: template,
	}

	return &res
}

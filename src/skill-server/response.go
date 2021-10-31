package main

import "fmt"

type SkillResponseType struct {
	Version  string       `json:"version"`
	Template TemplateType `json:"template,omitempty"`
}

type TemplateType struct {
	Outputs []OutputsType `json:"outputs,omitempty"`
}

type OutputsType struct {
	// SimpleText SimpleTextType `json:"simpleText,omitempty"`
	ItemCard ItemCardType `json:"itemCard,omitempty"`
	// Carousel CarouselType `json:"carousel"`
}

type SimpleTextType struct {
	Text string `json:"text"`
}

type ItemCardType struct {
	Profile  ICProfile `json:"profile"`
	ItemList []ICItem  `json:"itemList"`
}

type ICProfile struct {
	Title string `json:"title"`
}

type ICItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CarouselType struct {
	Type  string         `json:"type"`
	Items []ItemCardType `json:"items"`
}

func GetOneMenuReasponse(menu []string) *SkillResponseType {
	itemList := []ICItem{}
	for i, v := range menu {
		itemList = append(itemList, ICItem{Title: fmt.Sprintf("%d", i+1), Description: v})
	}
	itemCard := ItemCardType{
		Profile: ICProfile{
			Title: "점심",
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

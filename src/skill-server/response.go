package main

type SkillResponseType struct {
	Version  string       `json:"version"`
	Template TemplateType `json:"template,omitempty"`
}

type TemplateType struct {
	Outputs []OutputsType `json:"outputs,omitempty"`
}

type OutputsType struct {
	ItemCard ItemCardType `json:"itemCard,omitempty"`
}

type SimpleTextType struct {
	Text string `json:"text"`
}

type ItemCardType struct {
	Head     ICHead   `json:"head"`
	ItemList []ICItem `json:"itemList"`
}

type ICHead struct {
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

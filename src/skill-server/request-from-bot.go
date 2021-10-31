package main

type BotRequestType struct {
	Action ActionType `json:"action"`
}

type ActionType struct {
	Params map[string]string `json:"params"`
}

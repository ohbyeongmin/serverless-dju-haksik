package main

import (
	"time"

	"github.com/ohbyeongmin/daejeon-haksik/constants"
	"github.com/ohbyeongmin/daejeon-haksik/menu"
)

type SkillServerService interface {
	GetMenu(which constants.LunOrDin, weekday time.Weekday) []string
}

var HRCService SkillServerService = menu.HRCMenuService{}

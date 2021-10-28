package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/ohbyeongmin/daejeon-haksik/constants"
	"github.com/ohbyeongmin/daejeon-haksik/utils"
	"github.com/xuri/excelize/v2"
)

// func getToday() time.Weekday {
// 	t := time.Now()
// 	return t.Weekday()
// }

// func getTomorrow() time.Weekday {
// 	t := time.Now()
// 	return t.Add(time.Hour * 24).Weekday()
// }

type HRCMenuService struct{}

func (HRCMenuService) GetMenu(which constants.LunOrDin, weekday time.Weekday) []string {
	return mem.GetOne(which, weekday)
}

const (
	minRowNum int = 4
	maxRowNum int = 23
	minColNum int = 2
	maxColNum int = 7
)

type menu map[time.Weekday][]string

type menutable struct {
	table map[constants.LunOrDin]menu
}

var mem menutable

func InitMenu() {
	mem = menutable{
		table: make(map[constants.LunOrDin]menu),
	}
	mem.table[constants.LUNCH] = make(menu)
	mem.table[constants.DINNER] = make(menu)

	mem.parseMenuFile()
}

func (m *menutable) parseMenuFile() {
	f, err := excelize.OpenFile("diet.xlsx")
	utils.HandleErr(err)
	sheetName := f.GetSheetList()[0]

	rows, err := f.GetRows(sheetName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lunchOrDinner := constants.LUNCH
	for i, row := range rows {
		if i < minRowNum {
			continue
		}
		if i > maxRowNum {
			break
		}
		for j, colCell := range row {
			if j < minColNum || j >= maxColNum {
				if colCell == "석  식" {
					lunchOrDinner = constants.DINNER
				}
				continue
			}
			mem.table[lunchOrDinner][time.Weekday(j-1)] = append(mem.table[lunchOrDinner][time.Weekday(j-1)], colCell)
		}
	}
}

func (m menutable) GetOne(which constants.LunOrDin, weekDay time.Weekday) []string {
	return m.table[which][weekDay]
}

func (m menutable) GetAll(which constants.LunOrDin) [][]string {
	var allMenu = make([][]string, 5)
	for k, v := range m.table[which] {
		switch k {
		case time.Monday:
			allMenu[0] = append(allMenu[0], v...)
		case time.Tuesday:
			allMenu[1] = append(allMenu[1], v...)
		case time.Wednesday:
			allMenu[2] = append(allMenu[1], v...)
		case time.Thursday:
			allMenu[3] = append(allMenu[1], v...)
		case time.Friday:
			allMenu[4] = append(allMenu[1], v...)
		}
	}
	return allMenu
}

func Menu() *menutable {
	return &mem
}

func WriteFile() {
	f, _ := os.Create("test")
	defer f.Close()
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(mem.table)
	f.Write(buffer.Bytes())
	// fmt.Println(buffer.Bytes())
}

func ReadFile() {
	f, _ := os.Open("test")
	defer f.Close()
	var m menutable = menutable{}
	var buffer bytes.Buffer
	dec := gob.NewDecoder(&buffer)
	buffer.ReadFrom(f)
	dec.Decode(&m.table)
	fmt.Println(m.table[constants.LUNCH][time.Friday])
}

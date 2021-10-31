package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type LunOrDin int

const (
	LUNCH LunOrDin = iota
	DINNER
)

func HandleErr(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}

func isVaildCheckDay(target int) bool {
	if target < 1 || target > 31 {
		return false
	}
	return true
}

func CheckThisWeek(target int) bool {
	if !isVaildCheckDay(target) {
		return false
	}
	t := time.Now()
	if target < t.Day()-int(t.Weekday()) || target >= t.Day()+(7-int(t.Weekday())) {
		return false
	}

	return true
}

// 일(day) 을 요일로 변환
func DayToWeekday(target int) time.Weekday {
	if !isVaildCheckDay(target) {
		return -1
	}
	t := time.Now()
	return time.Weekday((int(t.Weekday()) + ((target - t.Day()) % 7)) % 7)
}

func StringToLunOrDin(s string) (LunOrDin, error) {
	if strings.Contains(s, "lunch") {
		return LUNCH, nil
	} else if strings.Contains(s, "dinner") {
		return DINNER, nil
	}
	return -1, errors.New("wrong input")
}

// day 가 null 일 경우 0 반환
// dateTag 가 null 일 경우 빈 문자열 반환
func ParseSysdate(sysdate string) (string, int) {
	var dateTag string
	var day string
	split := strings.Split(sysdate, ",")
	for _, v := range split {
		if strings.Contains(v, "dateTag") {
			dateTag = strings.Split(v, ":")[1]
		}
		if strings.Contains(v, "day") {
			day = strings.Split(v, ":")[1]
		}
	}
	dateTag = strings.Trim(strings.Trim(dateTag, " "), "\"")
	day = strings.Trim(strings.Trim(strings.Trim(day, " "), "}"), "\"")

	if dateTag == "null" {
		dateTag = ""
	}

	nDay, err := strconv.Atoi(day)
	if err != nil {
		nDay = 0
	}

	return dateTag, nDay
}

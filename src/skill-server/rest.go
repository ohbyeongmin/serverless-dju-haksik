package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ohbyeongmin/daejeon-haksik/utils"
)

const port string = ":3000"

func TodayHandler(rw http.ResponseWriter, r *http.Request) {
	var weekDay time.Weekday
	params := BotRequestType{
		Action: ActionType{
			Params: make(map[string]string),
		},
	}
	json.NewDecoder(r.Body).Decode(&params)

	dateTag, day := utils.ParseSysdate(params.Action.Params["sys_date"])
	lunchOrDinner, err := utils.StringToLunOrDin(params.Action.Params["lunch_dinner"])
	utils.HandleErr(err)

	// 이번주 체크
	// if !utils.CheckThisWeek(day) {
	// 	fmt.Println("unvalid day")
	// 	return
	// }
	t := time.Now()
	if day != 0 {
		weekDay = utils.DayToWeekday(day)
	} else if dateTag == "today" {
		weekDay = t.Weekday()
	} else if dateTag == "tomorrow" {
		weekDay = t.Add(time.Hour * 24).Weekday()
	} else if dateTag == "Monday" {
		weekDay = time.Monday
	} else if dateTag == "Tuesday" {
		weekDay = time.Tuesday
	} else if dateTag == "Wednesday" {
		weekDay = time.Wednesday
	} else if dateTag == "Thursday" {
		weekDay = time.Thursday
	} else if dateTag == "Friday" {
		weekDay = time.Friday
	}

	fmt.Println(lunchOrDinner, weekDay)
	menu := HRCService.GetMenu(lunchOrDinner, weekDay)
	json.NewEncoder(rw).Encode(GetOneMenuReasponse(menu))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func ServerStart() {
	r := mux.NewRouter()
	r.Use(jsonContentTypeMiddleware)
	r.HandleFunc("/today", TodayHandler).Methods("POST")

	fmt.Printf("Listen on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		utils.HandleErr(err)
	}
}

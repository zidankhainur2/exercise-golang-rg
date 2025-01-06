package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	var hour, minute int

	switch t := time.(type) {
	case string:
		parts := strings.Split(t, ":")
		if len(parts) != 2 {
			return "Invalid input"
		}
		h, err1 := strconv.Atoi(parts[0])
		m, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return "Invalid input"
		}
		hour, minute = h, m
	case []int:
		if len(t) != 2 {
			return "Invalid input"
		}
		hour, minute = t[0], t[1]
	case map[string]int:
		h, ok1 := t["hour"]
		m, ok2 := t["minute"]
		if !ok1 || !ok2 {
			return "Invalid input"
		}
		hour, minute = h, m
	case Time:
		hour, minute = t.Hour, t.Minute
	default:
		return "Invalid input"
	}

	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return "Invalid input"
	}

	period := "AM"
	if hour >= 12 {
		period = "PM"
	}
	if hour == 0 {
		hour = 12
	} else if hour > 12 {
		hour -= 12
	}

	return fmt.Sprintf("%02d:%02d %s", hour, minute, period)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}

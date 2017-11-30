package hub

import (
	_ "fmt"
	"log"
	"strconv"
	"time"
)

func TimeToStartTime(toConvert time.Time) time.Time {
	return time.Date(toConvert.Year(), toConvert.Month(), toConvert.Day(), 0, 0, 1, 0, toConvert.Location())
}

func TimeToEndTime(toConvert time.Time) time.Time {
	return time.Date(toConvert.Year(), toConvert.Month(), toConvert.Day(), 23, 59, 59, 0, toConvert.Location())
}

func FormatTimeHourMinute(toFormat time.Time) string {
	//	toFormat = RoundTime(toFormat)
	formatted := ""
	if toFormat.Hour() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Hour())
	} else {
		formatted += strconv.Itoa(toFormat.Hour()) + ""
	}
	formatted += ":"
	if toFormat.Minute() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Minute())
	} else {
		formatted += strconv.Itoa(toFormat.Minute())
	}
	return formatted
}

func FormatFullTime(toFormat time.Time) string {
	//	toFormat = RoundTime(toFormat)
	formatted := ""
	if toFormat.Hour() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Hour())
	} else {
		formatted += strconv.Itoa(toFormat.Hour()) + ""
	}
	formatted += ":"
	if toFormat.Minute() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Minute())
	} else {
		formatted += strconv.Itoa(toFormat.Minute())
	}
	formatted += "-"
	if toFormat.Day() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Day())
	} else {
		formatted += strconv.Itoa(toFormat.Day())
	}
	formatted += "-" + toFormat.Month().String() + "-" + strconv.Itoa(toFormat.Year())
	return formatted
}

func RoundTime(toRound time.Time) time.Time {
	toRoundMinuteString := ""
	if toRound.Minute() < 10 {
		toRoundMinuteString += "0" + strconv.Itoa(toRound.Minute())
	} else {
		toRoundMinuteString += strconv.Itoa(toRound.Minute())
	}
	minute, err := strconv.Atoi(toRoundMinuteString[1:2])
	var finalMinute string
	var finalHour string
	if err != nil {
		log.Fatal(err)
	}
	if minute == 0 {
		return toRound
	}
	if minute <= 2 {
		finalMinute = toRoundMinuteString[0:1] + "0"
		finalHour = strconv.Itoa(toRound.Hour())
	} else if minute > 2 && minute <= 7 {
		finalMinute = toRoundMinuteString[0:1] + "5"
		finalHour = strconv.Itoa(toRound.Hour())
	} else {
		newFirstMin, err := strconv.Atoi(toRoundMinuteString[0:1])
		if err != nil {
			log.Fatal(err)
		}
		if newFirstMin < 5 {
			newFirstMin++
			finalHour = strconv.Itoa(toRound.Hour())
		} else {
			if toRound.Hour() < 23 {
				toRound = toRound.Add(time.Hour * time.Duration(1))
				finalHour = strconv.Itoa(toRound.Hour())
			} else {
				toRound = toRound.Add(time.Hour * 24)
				finalHour = "0"
			}
			newFirstMin = 0
		}
		finalMinute = strconv.Itoa(newFirstMin) + "0"
	}
	finalMinuteInt, err := strconv.Atoi(finalMinute)
	if err != nil {
		log.Fatal()
	}
	finalHourInt, err := strconv.Atoi(finalHour)
	if err != nil {
		log.Fatal()
	}
	finalTime := time.Date(toRound.Year(), toRound.Month(), toRound.Day(), finalHourInt, finalMinuteInt, 0, 0, toRound.Location())
	return finalTime
}

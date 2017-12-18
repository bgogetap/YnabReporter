package report

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Month struct {
	month    string
	year     string
	networth int64
}

func newMonth(line string) Month {
	month, year := getDate(line)
	amount := getAmount(line)
	return Month{month, year, amount}
}

func getExistingMonth(monthList []Month, monthString string, yearString string) *Month {
	for index, month := range monthList {
		if month.month == monthString && month.year == yearString {
			return &monthList[index]
		}
	}
	return &Month{}
}

func addToMonth(month *Month, amount int64) {
	amount = month.networth + amount
	month.networth = amount
}

func getAmount(line string) int64 {
	value := strings.Split(line, ",")
	rawIn := value[9]
	rawOut := value[8]
	rawIn = strings.Replace(rawIn, "$", "", 1)
	rawIn = strings.Replace(rawIn, ".", "", 1)
	rawOut = strings.Replace(rawOut, "$", "", 1)
	rawOut = strings.Replace(rawOut, ".", "", 1)

	inflow, _ := strconv.ParseInt(rawIn, 0, 64)
	outflow, _ := strconv.ParseInt(rawOut, 0, 64)
	return inflow - outflow
}

func getDate(line string) (string, string) {
	pieces := strings.Split(line, ",")
	return pieces[2][1:3], pieces[2][7:]
}

func ParseMonth(registerFile *os.File) {
	var monthList []Month
	scanner := bufio.NewScanner(registerFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Account") {
			continue
		}
		m, y := getDate(line)
		matchedMonth := getExistingMonth(monthList, m, y)
		if matchedMonth.month == "" {
			monthList = append(monthList, newMonth(line))
		} else {
			addToMonth(matchedMonth, getAmount(line))
		}
	}
	var networth int64
	for i := len(monthList) - 1; i >= 0; i-- {
		month := monthList[i]
		print(month.month + " " + month.year)
		print(" : $")
		networth += month.networth
		networthString := strconv.FormatInt(networth, 10)
		firstPart := networthString[:len(networthString)-2]
		secondPart := networthString[len(networthString)-2:]
		newString := firstPart + "." + secondPart
		print(newString)
		println()
	}
}

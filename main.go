package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"ynab_reporter/report"
)

type Category struct {
	name string
	spent float64
}

func newCategory(line string) Category {
	pieces := strings.Split(line, ",")
	amount := pieces[5]
	return Category{pieces[1], priceFromLine(amount)}
}

func priceFromLine(line string) float64 {
	amount := strings.Replace(line, "$", "", 1)
	floatAmount, _ := strconv.ParseFloat(amount, 32)
	return floatAmount
}

func addToCategory(category *Category, amount float64) {
	newAmount := category.spent + amount
	category.spent = newAmount
}

func containsCategory(categoryList []Category, newCategory Category) *Category {
	for index,item := range categoryList {
		if (item.name == newCategory.name) {
			return &categoryList[index]
		}
	}
	return &Category{}
}

func parseBudget(budgetFile *os.File) {
	var categoryList []Category
	scanner := bufio.NewScanner(budgetFile)
	for scanner.Scan() {
		extractedCategory := newCategory(scanner.Text())
		matchedCategory := containsCategory(categoryList, extractedCategory)
		if (matchedCategory.name != "") {
			addToCategory(matchedCategory, extractedCategory.spent)
		} else {
			categoryList = append(categoryList, extractedCategory)
		}
	}
	//for _,item := range categoryList {
	//	fmt.Println(item.name + " : " + strconv.FormatFloat(item.spent, 'f', 2, 64))
	//}
}

func main() {
	args := os.Args[1:]
	budget := args[0]
	register := args[1]

	budgetFile, _ := os.Open(budget)
	defer budgetFile.Close()
	registerFile, _ := os.Open(register)
	defer registerFile.Close()

	parseBudget(budgetFile)
	report.ParseMonth(registerFile)
}
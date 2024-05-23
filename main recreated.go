package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

var reader *bufio.Reader

func main() {
	fmt.Println("Starting Application")
	var user string
	reader = bufio.NewReader(os.Stdin)

	user, _ = getInput("User: ")
	user_master_file := loadUser(user)

	fmt.Println(user_master_file)

	menu(user)

}

func getInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func menu(user string) {
	var action string
	action, _ = getInput(fmt.Sprintf(
		"What would you like to do, %v? \n"+
			"\t1. Enter Expense\n"+
			"\t2. View Accounts\n"+
			"\t3. Quit\n", user))
	switch action {
	case "1":
		menu(user)
	case "2":
		menu(user)
	case "3":
		break

	}

}

func loadUser(user string) *excelize.File {
	file, err := excelize.OpenFile(fmt.Sprintf("./Ledgers/%v_Master.xlsx", user))
	// file, err := excelize.OpenFile(fmt.Sprintf("Andre_Master.xlsx"))
	if err != nil {
		log.Fatal(err)
	}

	account_number, _ := file.GetCellValue("Sheet1", "A2")

	fmt.Printf("loaded user %v, with account number %v\n", user, account_number)
	return file
}

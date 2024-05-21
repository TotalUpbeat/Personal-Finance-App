package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("Starting Application")
	var user string
	var action int8

	fmt.Print("User: ")
	fmt.Scan(&user)
	user_master_file := loadUser(user)

	fmt.Println(user_master_file)

	for {
		fmt.Printf("What would you like to do, %v?\n", user)
		fmt.Println("\t1. Enter Expense")
		fmt.Println("\t2. View Accounts")
		fmt.Println("\t3. Quit")
		fmt.Scan(&action)

		if action == 3 {
			fmt.Printf("Thank you, %v", user)
			break
		}
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

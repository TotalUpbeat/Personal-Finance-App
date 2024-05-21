package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("Starting Application")
	var user string

	fmt.Print("User: ")
	fmt.Scan(&user)
	loadUser(user)
}

func loadUser(user string) {
	file, err := excelize.OpenFile(fmt.Sprintf("./Ledgers/%v_Master.xlsx", user))
	// file, err := excelize.OpenFile(fmt.Sprintf("Andre_Master.xlsx"))
	if err != nil {
		log.Fatal(err)
	}

	account_number, _ := file.GetCellValue("Sheet1", "A2")

	fmt.Printf("loaded user %v, with account number %v", user, account_number)
}

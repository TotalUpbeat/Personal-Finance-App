package main

import (
	"github.com/xuri/excelize/v2"
)

type ledger struct {
	ledger_name  string
	ledger_file  *excelize.File
	start_amount float64
	end_amount   float64
}

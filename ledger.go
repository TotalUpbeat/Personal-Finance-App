package main

import (
	"github.com/xuri/excelize/v2"
)

type ledger struct {
	ledger_name     string
	ledger_file     *excelize.File
	start_amount    float64
	end_amount      float64
	journal_entries *journal_entry
}

type journal_entry struct {
	date        string
	description string
	credit      float64
	debit       float64
}

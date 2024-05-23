package main

import (
	"github.com/xuri/excelize/v2"
)

type Ledger struct {
	ledger_name     string
	ledger_file     *excelize.File
	start_amount    float64
	end_amount      float64
	journal_entries map[string][]Journal_Entry
}

type Journal_Entry struct {
	date        string
	description string
	credit      float64
	debit       float64
}

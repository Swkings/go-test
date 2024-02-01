package test

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestExcel(t *testing.T) {
	excelWrite()
	excelRead()
}

func excelRead() {
	f, err := excelize.OpenFile("../resource/data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	sheets := f.GetSheetList()
	// Get all the rows in the Sheet1.
	for _, sheet := range sheets {
		fmt.Print(sheet, "\n")
		rows, _ := f.GetRows(sheet)
		for _, row := range rows {
			for _, colCell := range row {
				if colCell == "" {
					fmt.Print("/", "\t")
				} else {
					fmt.Print(colCell, "\t")
				}
			}
			fmt.Println()
		}
	}
}

func excelWrite() {
	f := excelize.NewFile()
	// Create a new sheet.
	index, _ := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello")
	f.SetCellValue("Sheet2", "B2", "World")

	f.SetCellValue("Sheet1", "A1", 5)
	f.SetCellValue("Sheet1", "B1", 15)
	f.SetCellValue("Sheet1", "A2", 100)
	f.SetCellValue("Sheet1", "B2", 200)
	f.SetActiveSheet(index)
	f.SetActiveSheet(index - 1)
	// Save spreadsheet by the given path.
	err := f.SaveAs("../resource/data.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

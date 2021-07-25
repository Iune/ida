package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/Iune/ida/pkg/contest"
)

func SaveContestDetails(countriesFilePath string, entriesFilePath string, outputPathPrefix string) {
	countries := contest.LoadCountries(countriesFilePath)
	entries := contest.LoadEntries(entriesFilePath, countries)

	writeSpreadsheet(entries, outputPathPrefix)
	writeJSON(countries, entries, outputPathPrefix)
}

func writeSpreadsheet(entries []contest.Entry, outputPathPrefix string) {
	file := buildSpreadsheet(entries)
	file = setSpreadsheetStyles(file)
	if err := file.SaveAs(fmt.Sprintf("%s.xlsx", outputPathPrefix)); err != nil {
		fmt.Println(err)
	}
}

func buildSpreadsheet(entries []contest.Entry) *excelize.File {
	file := excelize.NewFile()
	sheet := "Sheet1"

	// Header row
	headerRow := []interface{}{"Draw", "Country", "Flag", "Artist", "Song", "Total", "Count"}
	for _, entry := range entries {
		name, _ := entry.Country.PrimaryName()
		headerRow = append(headerRow, name)
	}
	file.SetSheetRow(sheet, "A1", &headerRow)

	// Entry rows
	for i, entry := range entries {
		currentRow := i + 2

		// Entry details
		name, _ := entry.Country.PrimaryName()
		entryRow := []interface{}{i + 1, name, entry.Flag(), entry.Artist, entry.Song}

		rowCell, _ := excelize.CoordinatesToCellName(1, currentRow)
		file.SetSheetRow(sheet, rowCell, &entryRow)

		// Formulas
		startCell, _ := excelize.CoordinatesToCellName(8, currentRow)
		endCell, _ := excelize.CoordinatesToCellName(8+len(entries), currentRow)

		sumCell, _ := excelize.CoordinatesToCellName(6, currentRow)
		countCell, _ := excelize.CoordinatesToCellName(7, currentRow)
		file.SetCellFormula(sheet, sumCell, fmt.Sprintf("=SUM(%s,%s)", startCell, endCell))
		file.SetCellFormula(sheet, countCell, fmt.Sprintf("=COUNT(%s,%s)", startCell, endCell))

		// X for voter
		voterCell, _ := excelize.CoordinatesToCellName(8+i, currentRow)
		file.SetCellValue(sheet, voterCell, "X")
	}

	return file
}

func setSpreadsheetStyles(file *excelize.File) *excelize.File {
	sheet := "Sheet1"
	headerStyle, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Font:      &excelize.Font{Family: "Palatino", Size: 12, Bold: true},
	})

	headerRotatedStyle, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", TextRotation: 90},
		Font:      &excelize.Font{Family: "Palatino", Size: 12, Bold: true},
	})

	defaultStyle, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Font:      &excelize.Font{Family: "Palatino", Size: 12},
	})

	rows, _ := file.GetRows(sheet)
	cols, _ := file.GetCols(sheet)

	// Header row, entry details
	for i := 1; i < 8; i++ {
		cell, _ := excelize.CoordinatesToCellName(i, 1)
		file.SetCellStyle(sheet, cell, cell, headerStyle)
	}

	// Header row, vote details
	for i := 6; i <= len(cols); i++ {
		cell, _ := excelize.CoordinatesToCellName(i, 1)
		file.SetCellStyle(sheet, cell, cell, headerRotatedStyle)
	}

	// Entry rows
	for i := 2; i <= len(rows); i++ {
		for j := 1; j <= len(cols); j++ {
			cell, _ := excelize.CoordinatesToCellName(j, i)
			file.SetCellStyle(sheet, cell, cell, defaultStyle)
		}
	}

	return file
}

func writeJSON(countries []contest.Country, entries []contest.Entry, outputPathPrefix string) {
	contest := contest.Contest{
		Entries:   entries,
		Countries: countries,
		Voters:    countries,
	}

	data, _ := json.MarshalIndent(&contest, "", "\t")
	if err := os.WriteFile(fmt.Sprintf("%s.json", outputPathPrefix), data, 0666); err != nil {
		fmt.Println(err)
	}
}

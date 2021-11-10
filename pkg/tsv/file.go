package tsv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/zguillez/go-tools/files"
	"github.com/zguillez/go-tools/system"
)

func Matrix(path string) (data [][]string) {
	rows := Rows(path)
	for _, row := range rows {
		items := Columns(row)
		data = append(data, items)

	}
	return data
}

func Rows(path string) (rows [][]string) {
	f, err := os.Open(path)
	system.CheckError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	rows, err = reader.ReadAll()
	system.CheckError(err)

	return rows
}
func Columns(row []string) (columns []string) {
	for i, item := range row {
		if i != len(row)-1 {
			columns = append(columns, item)
		}
	}

	return columns
}

func SaveFile(filePath, text string) {
	files.SaveFile(filePath, text)
}

func Stringify(matrix [][]string) (text string) {
	text = ""
	for _, row := range matrix {
		for _, column := range row {
			text = fmt.Sprintf("%s%s\t", text, column)
		}
		text = fmt.Sprintf("%s\n", text)
	}

	return text
}

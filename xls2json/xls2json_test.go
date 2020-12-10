package xls2json

import "testing"

// go test ./ -v -test.run=Test_ConvertFirstSheet
func Test_ConvertFirstSheet(t *testing.T) {
	result, err := ConvertFirstSheet("./sample.xlsx", DefaultOptions)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(result.OnlyJson())
	t.Log(result.Summary())
}

// go test ./ -v -test.run=Test_ConvertOneSheet
func Test_ConvertOneSheet(t *testing.T) {
	result, err := ConvertOneSheet("./sample.xlsx", DefaultSheetName, DefaultOptions)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(result.OnlyJson())
	t.Log(result.Summary())
}

// go test ./ -v -test.run=Test_ConvertAllSheets
func Test_ConvertAllSheets(t *testing.T) {
	result, err := ConvertAllSheets("./sample.xlsx", DefaultOptions)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(result.OnlyJson())
	t.Log(result.Summary())
}

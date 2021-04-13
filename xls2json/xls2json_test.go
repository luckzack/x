package xls2json

import "testing"

// go test ./ -v -test.run=Test_ConvertFirstSheet -count=1
func Test_ConvertFirstSheet(t *testing.T) {
	opt := Options{
		HandleEmptyRow: 1,
	}

	result, err := ConvertFirstSheet("./sample.xlsx", &opt)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(result.OnlyJson())
	t.Log(result.Summary())
}

// go test ./ -v -test.run=Test_ConvertOneSheet -count=1
func Test_ConvertOneSheet(t *testing.T) {
	result, err := ConvertOneSheet("./sample.xlsx", DefaultSheetName, DefaultOptions)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(result.OnlyJson())
	t.Log(result.Summary())
}

// go test ./ -v -test.run=Test_ConvertAllSheets -count=1
func Test_ConvertAllSheets(t *testing.T) {
	result, err := ConvertAllSheets("./sample.xlsx", DefaultOptions)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(result.OnlyJson())
	t.Log(result.Summary())
}

package xls2json

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

type sheetJSONMap []map[string]interface{}

type Options struct {
	// 从表格的第几行开始，默认0，即第一行
	StartRow int
	// 自定义表头，对应json中的key，默认为空取表格的第一行
	CustomHeaders []string
	// 跳过表格的第多少行，从0开始
	SkipRows []int
	// 只解析表格中最多多少行，默认无上限
	MaxRows int
	// 遇到多少个空行怎么处理
	HandleEmptyRow int //0-默认写入空行，1-skip跳过空行，2-退出

}

var DefaultOptions = &Options{
	StartRow:           0,
	HandleEmptyRow: 0,
}
var DefaultSheetName = "Sheet1"

///////////////////
type TotalResult struct {
	SheetResults map[string]*SheetResult //key为sheetName
	Bytes        []byte
}

func (tr *TotalResult) Summary() string {

	sheetSummary := make(map[string]string)
	for _, sheetResult := range tr.SheetResults {
		sheetSummary[sheetResult.SheetName] = sheetResult.Summary()
	}

	m := map[string]interface{}{
		"SheetNum":     len(tr.SheetResults),
		"SheetSummary": sheetSummary,
	}
	bytes, _ := json.Marshal(&m)
	return string(bytes)
}

func (tr *TotalResult) OnlyJsonBytes() []byte {
	m := make(map[string]interface{})
	for _, sr := range tr.SheetResults {
		m[sr.SheetName] = sr.SheetJson
	}

	jsonBytes, _ := json.MarshalIndent(&m, "", "	")

	return jsonBytes
}

func (tr *TotalResult) OnlyJson() string {
	return string(tr.OnlyJsonBytes())
}

type SheetResult struct {
	SheetName string
	RawBytes  []byte
	Rows      int
	CostTime  time.Duration
	SheetJson sheetJSONMap
}

func (sr *SheetResult) Summary() string {
	m := map[string]interface{}{
		"SheetName": sr.SheetName,
		"Rows":      sr.Rows,
		"CostTime":  sr.CostTime.String(),
		"Size":      len(sr.RawBytes),
	}
	bytes, _ := json.Marshal(&m)
	return string(bytes)
}

func (sr *SheetResult) OnlyJson() string {
	return string(sr.RawBytes)
}

///////////////////

func parseSheetDataToJson(sheetName string, sheetData *xlsx.Sheet, options *Options) (*SheetResult, error) {
	if options == nil {
		options = DefaultOptions
	}

	startTime := time.Now()

	info := make(sheetJSONMap, 0)

	markMap := make(map[int]string)
	setHeader := true

	for rowIndex, row := range sheetData.Rows {

		if rowIndex < options.StartRow {
			continue
		}

		skip := false
		for _, r := range options.SkipRows {
			if rowIndex == r {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		var jsonMap map[string]interface{}
		if rowIndex != 0 {
			jsonMap = make(map[string]interface{})
		}

		fmt.Printf("row idx: %d\n", rowIndex)
		if len(row.Cells)<1 {
			if options.HandleEmptyRow == 1 {
				continue
			}else if options.HandleEmptyRow == 2 {
				break
			}
		}

		for cellIndex, cell := range row.Cells {
			fmt.Printf("cell idx: %d, val: %s \n", cellIndex, cell.Value)

			if len(options.CustomHeaders) > 0 && cellIndex >= len(options.CustomHeaders) {
				break
			}
			cell.Value = strings.TrimSpace(cell.Value)

			if setHeader {
				if len(options.CustomHeaders) > 0 {
					// 自定义表头
					markMap[cellIndex] = options.CustomHeaders[cellIndex]
				} else {
					markMap[cellIndex] = cell.Value
				}

			} else {
				if jsonMap != nil {
					jsonMap[markMap[cellIndex]] = cell.Value
				}
			}
		}
		setHeader = false

		if jsonMap != nil {
			info = append(info, jsonMap)
		}

		if options.MaxRows > 0 && rowIndex >= options.MaxRows {
			break
		}

	}

	jsonBytes, err := json.MarshalIndent(&info, "", "	")
	if err != nil {
		return nil, err
	}

	return &SheetResult{
		SheetName: sheetName,
		CostTime:  time.Since(startTime),
		RawBytes:  jsonBytes,
		Rows:      len(info),
		SheetJson: info,
	}, nil
}

// Convert first excel sheet to json
func ConvertFirstSheet(xlsxPath string, options *Options) (*SheetResult, error) {
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		return nil, err
	}
	sheetData, has := xlsxFile.Sheet[DefaultSheetName]

	if !has {
		return nil, fmt.Errorf("no such sheet: %s", DefaultSheetName)
	}

	return parseSheetDataToJson(DefaultSheetName, sheetData, options)
}

// Convert specified one excel sheet to json
func ConvertOneSheet(xlsxPath string, sheetName string, options *Options) (*SheetResult, error) {
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		return nil, err
	}
	sheetData, has := xlsxFile.Sheet[sheetName]

	if !has {
		return nil, fmt.Errorf("no such sheet: %s", sheetName)
	}

	return parseSheetDataToJson(sheetName, sheetData, options)
}

// Convert all excel sheets to json
func ConvertAllSheets(xlsxPath string, options *Options) (*TotalResult, error) {
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		return nil, err
	}

	tr := &TotalResult{
		SheetResults: make(map[string]*SheetResult),
	}

	for _, sheetData := range xlsxFile.Sheets {
		sheetResult, err := parseSheetDataToJson(sheetData.Name, sheetData, options)
		if err != nil {
			return nil, fmt.Errorf("parse sheel='%s' fail: %s",
				sheetData.Name, err.Error())
		}

		tr.SheetResults[sheetData.Name] = sheetResult
	}
	return tr, nil
}

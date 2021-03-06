// Copyright 2020 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"context"
	"database/sql/driver"
	"go/types"
	"regexp"
	"strings"

	"github.com/ruslanBik4/dbEngine/dbEngine"

	"github.com/ruslanBik4/httpgo/logs"
	"github.com/ruslanBik4/httpgo/typesExt"
)

type ColumnDecor struct {
	dbEngine.Column
	IsHidden, IsReadOnly, IsSlice bool
	InputType                     string
	SelectOptions                 map[string]string
	PatternList                   dbEngine.Table
	PatternName                   string
	PlaceHolder                   string
	Label                         string
	pattern                       string
	patternDesc                   string
	Value                         interface{}
}

var regPattern = regexp.MustCompile(`\{"pattern":\s*"([^"]+)"\}`)

func NewColumnDecor(column dbEngine.Column, patternList dbEngine.Table) *ColumnDecor {

	colDec := &ColumnDecor{Column: column, PatternList: patternList}
	comment := colDec.Comment()
	if m := regPattern.FindAllStringSubmatch(comment, -1); len(m) > 0 {
		colDec.getPattern(m[0][1])
		colDec.Label = regPattern.ReplaceAllString(comment, colDec.patternDesc)
	} else if column.Comment() > "" {
		colDec.Label = column.Comment()
	} else {
		colDec.Label = column.Name()
	}

	colDec.InputType = colDec.inputType()

	return colDec
}

func (col *ColumnDecor) Placeholder() string {
	if col.PlaceHolder > "" {
		return col.PlaceHolder
	}

	if col.pattern > "" {
		return col.pattern
	}

	return col.Label
}

func (col *ColumnDecor) Pattern() string {
	if col.pattern > "" {
		return col.pattern
	}

	if name := col.PatternName; name > "" {
		col.getPattern(name)
	} else if col.BasicTypeInfo() == types.IsInteger {
		col.pattern = `[0-9]+`
	} else if col.BasicTypeInfo() == types.IsFloat {
		col.pattern = `[+-]?\d+(\.\d{2})?`
	} else if col.BasicTypeInfo() == types.IsComplex {
		col.pattern = `[+-]?\d+(\.\d{2})?`
	}

	return col.pattern
}

func (col *ColumnDecor) GetFields(columns []dbEngine.Column) []interface{} {
	if len(columns) == 0 {
		return []interface{}{&col.Value}
	}

	v := make([]interface{}, len(columns))
	for i, c := range columns {
		switch c.Name() {
		case "pattern":
			v[i] = &col.pattern
		case "description":
			v[i] = &col.patternDesc
		}
	}

	return v
}

var regName = regexp.MustCompile(`\w+`)

func (col *ColumnDecor) getPattern(name string) {
	if col.PatternList != nil && !regName.MatchString(name) {
		err := col.PatternList.SelectAndScanEach(context.Background(),
			func() error {
				return nil
			},
			col,
			dbEngine.ColumnsForSelect("pattern", "description"),
			dbEngine.WhereForSelect("name"),
			dbEngine.ArgsForSelect(name),
		)
		if err != nil {
			logs.ErrorLog(err, "")
		}
	}

	if col.pattern == "" {
		col.pattern = name
	}
}

func (col *ColumnDecor) Type() string {
	const email = "email"
	const tel = "phone"

	if strings.HasPrefix(col.Name(), email) {
		return email
	} else if strings.HasPrefix(col.Name(), tel) {
		return "tel"
	}

	return col.Column.Type()
}

func (col *ColumnDecor) GetValues() (values []interface{}) {

	switch val := col.Value.(type) {
	case []interface{}:
		values = val
	case []string:
		values = make([]interface{}, len(val))
		for i, val := range val {
			values[i] = val
		}
		col.IsSlice = true

	case []int32:
		values = make([]interface{}, len(val))
		for i, val := range val {
			values[i] = val
		}
		col.IsSlice = true

	case []int64:
		values = make([]interface{}, len(val))
		for i, val := range val {
			values[i] = val
		}
		col.IsSlice = true

	case []float32:
		values = make([]interface{}, len(val))
		for i, val := range val {
			values[i] = val
		}
		col.IsSlice = true

	case []float64:
		values = make([]interface{}, len(val))
		for i, val := range val {
			values[i] = val
		}
		col.IsSlice = true

	case nil:
		if d := col.Default(); d > "" {
			values = append(values, d)
		}
	case driver.Valuer:
		v, err := val.Value()
		if err != nil {
			logs.ErrorLog(err, "val.Value")
		} else {
			values = append(values, v)
		}
	default:
		values = append(values, val)
	}

	if len(values) == 0 {
		values = append(values, nil)
	}

	return
}

func (col *ColumnDecor) InputName(i int) string {
	if col.IsSlice {
		return col.Name() + "[]"
	}

	return col.Name()
}

func (col *ColumnDecor) inputType() string {
	if col.IsHidden {
		return "hidden"
	}

	switch col.Type() {
	case "date", "_date":
		return "date"
	case "datetime", "datetimez", "timestampt", "timestamptz", "time", "_timestampt", "_timestamptz", "_time":
		return "datetime"
	case "email", "tel", "password", "url":
		return col.Type()
	default:
		if typesExt.IsNumeric(col.BasicTypeInfo()) {
			return "number"
		}
		if col.BasicType() == types.Bool {
			return "checkbox"
		}

		return "text"
	}
}

type Button struct {
	Title      string
	Position   bool
	ButtonType string
}

type BlockColumns struct {
	Buttons            []Button
	Columns            []*ColumnDecor
	Id                 int
	Multype            bool
	Title, Description string
}

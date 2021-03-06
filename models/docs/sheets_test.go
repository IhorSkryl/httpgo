// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package docs

import (
	"fmt"
	"testing"
)

const spreadsheetId = "1EvNM788L-CC7N1kYIieQZEuinpmI7yVzu_mV75DF3cM"

func TestReadGoogleSheets(t *testing.T) {
	var sheet SheetsGoogleDocs
	fileName := ""

	defer func() {
		err := recover()
		if err != nil {
			t.Error(err)
		}
	}()
	if err := sheet.Init(); err != nil {
		t.Errorf("error initialization: filename=%s, error=%q", fileName, err)
	}

	fmt.Printf("before read")
	readRange := "шаблон!A1:F10"
	if resp, err := sheet.Read(spreadsheetId, readRange); err != nil {
		t.Errorf("Error during reading sheet %v", err)
	} else {
		fmt.Printf("%v", resp)
		for idx, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			for idx := range row {
				fmt.Printf("%s | ", row[idx])
			}
			fmt.Printf(", %d\n", idx)
		}

	}

	t.Skipped()
}

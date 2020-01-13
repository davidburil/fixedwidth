/*
   Copyright 2019 The David Buril Cardozo de Oliveira

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// Package fixedwith propose automated infer the columns in fixed-width text file
package fixedwith

import (
	"bufio"
	"errors"
	"io"
	"regexp"
)

// InferColumnsIndex Infer Column Widths of a Fixed-Width Text File
func InferColumnsIndex(reader *bufio.Reader, sampleValue uint) (columnsIndex []uint, err error) {

	if reader == nil {
		return nil, errors.New("the reader parameter is required")
	}

	var columns []bool

	for i := uint(0); i < sampleValue || sampleValue == 0; i++ {

		line, _, err := reader.ReadLine()

		if err != nil && err != io.EOF {
			return nil, err
		}

		if err == io.EOF {
			break
		}

		if columns == nil {
			columns = make([]bool, len(line))
		}

		for len(columns) < len(line) {
			columns = append(columns, false)
		}

		index := findAllStringIndex(line)

		for _, value := range index {
			for j := value[0] + 1; j < value[1]; j++ {
				columns[j] = true
			}
		}
	}

	result := parseColumnIndex(columns)

	return result, nil

}

// parseColumnIndex parse column index
func parseColumnIndex(columns []bool) []uint {
	result := make([]uint, 0)

	beforeValue := false
	beforeKey := 0

	for key, value := range columns {

		if value && !beforeValue {
			result = append(result, uint(beforeKey))
		}

		beforeValue = value
		beforeKey = key

	}
	return result
}

// findAllStringIndex build regex search columns in text file
func findAllStringIndex(text []byte) [][]int {
	r, _ := regexp.Compile(`[^\s]+`)

	return r.FindAllStringIndex(string(text), -1)
}

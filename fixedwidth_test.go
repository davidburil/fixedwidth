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

package fixedwith

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestInferColumnsIndex(t *testing.T) {
	s := strings.NewReader(`String              Bool  Int Int8 Int16 Int32 Int64 Uint Uint8 Uint16 Uint32 Uint64 Float32 Float64 Time                 
Test String         true  -1  -2   -3    -4    -5    1    2     3      4      5      1.5     2.5     2017-12-27T13:48:03Z 
Another test string false 0   0    0     0     0     0    0     0      0      0      0       0       0001-01-01T00:00:00Z 
Another string test false 0   0    0     0     0     0    0     0      0      0      0       0       0001-01-01T00:00:00Z `)

	bufioReader := bufio.NewReader(s)
	columnsIndex, _ := InferColumnsIndex(bufioReader, 0)

	fmt.Println(columnsIndex)
}

func TestInferColumnsIndexAndPrintFile(t *testing.T) {
	file, _ := os.Open("resources/dados2.txt")

	reader := bufio.NewReader(file)

	columnsIndex, _ := InferColumnsIndex(reader, 10)

	file.Seek(0, 0)

	fmt.Println(columnsIndex)

	printFile(reader, columnsIndex, 10)

}

func printFile(reader *bufio.Reader, columnsIndex []uint, linesPrint uint) {

	if reader == nil {
		panic("the file parameter is required")
	}

	if columnsIndex == nil {
		panic("the columns parameter is required")
	}

	for i := uint(0); i < linesPrint || linesPrint == 0; i++ {
		line, _, err := reader.ReadLine()

		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}

		var t []string
		var v int

		for _, value := range columnsIndex {

			if value == 0 {
				v = 0
				continue
			}
			t = append(t, string(line[v:value]))
			v = int(value)
		}

		t = append(t, string(line[v:]))

		fmt.Println(t)

	}
}

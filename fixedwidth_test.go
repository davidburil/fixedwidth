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
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestInferColumnsIndex(t *testing.T) {
	s := strings.NewReader(`String              Bool  Int Int8 Int16 Int32 Int64 Uint Uint8 Uint16 Uint32 Uint64 Float32 Float64 Time                 
Test String         true  -1  -2   -3    -4    -5    1    2     3      4      5      1.5     2.5     2017-12-27T13:48:03Z 
Another test string false 0   0    0     0     0     0    0     0      0      0      0       0       0001-01-01T00:00:00Z 
Another string test false 0   0    0     0     0     0    0     0      0      0      0       0       0001-01-01T00:00:00Z `)

	bufioReader := bufio.NewReader(s)
	indexColumn, _ := InferColumnsIndex(bufioReader, 0)

	fmt.Println(indexColumn)
}

func TestReadFile(t *testing.T) {

	file, _ := os.Open("resources/dados2.txt")

	reader := bufio.NewReader(file)

	i, _ := InferColumnsIndex(reader, 10)

	readFile(file, i)

}

// readFile teste de leitura do arquivo.
func readFile(file *os.File, columns []uint) (map[int][]string, error) {

	var retorno map[int][]string

	if file == nil {
		return nil, errors.New("the file parameter is required")
	}

	if columns == nil {
		return nil, errors.New("the columns parameter is required")
	}

	retorno = make(map[int][]string)

	reader := bufio.NewReader(file)

	for i := 0; true; i++ {
		line, _, err := reader.ReadLine()

		if err != nil && err != io.EOF {
			return nil, err
		}

		if err == io.EOF {
			break
		}

		var t []string
		var v int

		for _, value := range columns {

			if value == 0 {
				v = 0
				continue
			}
			t = append(t, string(line[v:value]))
			v = int(value)
		}

		t = append(t, string(line[v:]))

		retorno[i] = t

	}

	return retorno, nil
}

func TestRegex(t *testing.T) {
	file, _ := os.Open("resources/dados2.txt")

	reader := bufio.NewReader(file)

	r, err := regexp.Compile(`[^\s]+`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	var columns []bool

	for i := uint(0); i < uint(3); i++ {
		line, _, _ := reader.ReadLine()

		index := r.FindAllStringIndex(string(line), -1)

		if columns == nil {
			columns = make([]bool, len(line))
		}

		for len(columns) < len(line) {
			columns = append(columns, false)
		}

		for _, value := range index {
			for j := value[0] + 1; j < value[1]; j++ {
				columns[j] = true
			}
		}

		fmt.Println(index)

	}

	result := make([]uint, 0)

	valorAnterior := false
	chaveAnterior := 0

	for key, value := range columns {

		if value && !valorAnterior {
			result = append(result, uint(chaveAnterior))
		}

		valorAnterior = value
		chaveAnterior = key

	}

	fmt.Println(columns)
	fmt.Println(result)
}

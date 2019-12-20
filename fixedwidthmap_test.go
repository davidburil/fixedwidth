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

package fixedwithmap

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"
)

func TestReadFile(t *testing.T) {

	file, _ := os.Open("resources/dados2.txt")

	i, _ := InferColumnsIndex(file, 200)

	readFile(file, i)

}

// readFile teste de leitura do arquivo.
func readFile(file *os.File, columns []int) (map[int][]string, error) {

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
			v = value
		}

		t = append(t, string(line[v:]))

		retorno[i] = t

	}

	return retorno, nil
}

func TestRegex(t *testing.T) {
	file, _ := os.Open("resources/dados2.txt")

	reader := bufio.NewReader(file)

	r, err := regexp.Compile(`.+?\s+`)

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
				columns[j] = columns[j] || isNonWhiteSpace(line[j]) || i == 0
			}
		}

	}

	result := make([]uint, 0)

	for key, value := range columns {
		if !value {
			result = append(result, uint(key))
		}
	}

	fmt.Println(columns)
	fmt.Println(result)
}

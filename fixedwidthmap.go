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
	"io"
	"os"
	"unicode"
)

// InferColumnsIndex Metodo publico para fazer a inferência o index das colunas.
func InferColumnsIndex(file *os.File, sampleValue uint) (columns []int, err error) {

	dataSampleFromFile, e := extractColumnArrayWhiteSpace(file, sampleValue)

	if e != nil {
		return nil, e
	}

	columns, e = inferColumnsIndexAlignLeft(dataSampleFromFile)

	return columns, e

}

// extractColumnArrayWhiteSpace Metodo responsável por extrair um array com o resultado do processamento da amostra.
func extractColumnArrayWhiteSpace(file *os.File, sampleValue uint) (columns []bool, err error) {

	if file == nil {
		return nil, errors.New("the file parameter is required")
	}

	reader := bufio.NewReader(file)

	for i := uint(0); i < sampleValue || sampleValue == 0; i++ {
		line, _, e := reader.ReadLine()

		if e != nil && e != io.EOF {
			return nil, e
		}

		if e == io.EOF {
			break
		}

		if columns == nil {
			columns = make([]bool, len(line))
		}

		for len(columns) < len(line) {
			columns = append(columns, false)
		}

		for key, value := range line {
			columns[key] = columns[key] || isNonWhiteSpace(value)
		}

	}

	rewindFile(file)

	return columns, nil
}

// rewindFile Metodo responsavel por voltar o ponteiro para o inicio do arquivo.
func rewindFile(file *os.File) {
	file.Seek(0, 0)
}

// inferColumnsIndexAlignLeft Metodo responsável por inferir as posicoes das colunas alinhadas a esquerda.
func inferColumnsIndexAlignLeft(data []bool) (columns []int, err error) {

	if data == nil || len(data) < 4 {
		return nil, errors.New("the data parameter is required and must have at least 4 elements")
	}

	beforeValue := false

	for key, value := range data {
		if value && !beforeValue {
			columns = append(columns, key)
		}
		beforeValue = value
	}

	if columns == nil {
		return nil, errors.New("unable to identify columns")
	}

	return columns, nil
}

// isNonWhiteSpace retorna true caso não seja um espaço em branco.
func isNonWhiteSpace(b byte) bool {
	return !unicode.IsSpace(rune(b))
}

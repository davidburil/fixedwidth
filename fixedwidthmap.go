package fixedwithmap

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

// InferColumnsIndex Metodo publico para fazer a inferência o index das colunas.
func InferColumnsIndex(fileName string, sampleValue int) {

	dataSampleFromFile := extractDataSamplefromFile(fileName, sampleValue)

	print(dataSampleFromFile)

}

// extractDataSamplefromFile Extrai os dados de amostra do arquivo convertido para uma matrix de boleano.
func extractDataSamplefromFile(fileName string, sampleValue int) map[int][]bool {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	dataSample := make(map[int][]bool)

	reader := bufio.NewReader(file)

	for i := 0; i < sampleValue; i++ {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		dataSample[i] = convertArrayByteToArrayBool(line)
	}
	return dataSample
}

// convertArrayByteToArrayBool converte um array de byte em um array de booleano.
func convertArrayByteToArrayBool(data []byte) []bool {

	var r []bool
	for _, b := range data {
		r = append(r, isNonWhiteSpace(b))
	}
	return r
}

// isNonWhiteSpace retorna true caso não seja um espaço em branco.
func isNonWhiteSpace(b byte) bool {
	return !unicode.IsSpace(rune(b))
}

package fixedwithmap

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {

	file, err := os.Open("resources/dados2.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	linha, _, _ := reader.ReadLine()

	bits := ByteToBit(linha)

	fmt.Println(bits)
}
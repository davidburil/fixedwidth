package fixedwithmap

import (
	"testing"
)

func TestReadFile(t *testing.T) {

	InferColumnsIndex("resources/dados2.txt", 200)

}

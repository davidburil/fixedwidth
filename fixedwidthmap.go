package fixedwithmap

// ByteToBit converte um array de byte em um array de booleano.
func ByteToBit(data []byte) (r []bool) {

	var bits []bool

	for _, b := range data{
		bits = append(bits, spaceZero(b))
	}

	return bits
}

// spaceZero retorna false se o byte for igual a um espaço em branco.
func spaceZero(b byte) bool{
	return b != 32
}
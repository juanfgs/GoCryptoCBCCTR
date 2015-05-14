package tools

// XORS 2 []byte
func XOR(a, b []byte) []byte {

	var result []byte = make([]byte, len(a))

	for i := 0; i < len(a); i++ {

		result[i] = a[i] ^ b[i]

	}

	return result

}

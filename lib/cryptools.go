package cryptools

func XOR(a,b string) string {
	aBytes := []byte(a)
	bBytes := []byte(b) 	
	var result []byte = make([]byte, len(aBytes))
	for idx, valA := range aBytes {
		result[idx] = valA ^ bBytes[idx]
	}
	return string(result)
}

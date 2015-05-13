package tools

import (


)



// XORS 2 []byte
func XOR(a, b []byte) []byte {
/*
	var dst []byte
	var src *[]byte

	if len(a) >= len(b) {
		pad := len(a ) - len(b)
		dst = make([]byte, pad)
		dst = append(dst, b...)
		src = &a
	} else {
		pad := len(b) - len(a)
		dst = make([]byte, pad)
		dst = append(dst, a...)		

		src = &b
	}
*/

	
	var result []byte= make([]byte, len(a))

	for i := 0; i < len(a);i++ {
	/*	if a[i] == 48 {
			a[i] = 0
		}
		if b[i] == 48 {
			b[i] = 0
		}	 */
		result[i] = a[i] ^ b[i]
			
	}

	return result


}


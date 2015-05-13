package cbc

import (
	"testing"

)

func TestEncrypt(t *testing.T){
	cipher := New([]byte("140b41b22a29beb4061bda66b6747e14"), 16)
	
	 cipher.Encrypt([]byte("foobar"))
/*	for _,val := range ct {
		fmt.Printf("%d,", val)
	} */
}

func TestDecrypt(t *testing.T) {

	for _,n := range DecryptTest {
		if val := string(n.cipher.Decrypt(n.input)); val != n.expected {
			t.Fatalf("Expected: %s \n Got: %s", n.expected, val)
		}
	}
}

func TestPad(t *testing.T) {

	for _,n := range PadTest {
		if  res := Pad(n.input,n.blockSize); len(res) % n.blockSize != 0 {
			t.Fatalf("Resulting padded string must be of the block size")
		}
	}


}

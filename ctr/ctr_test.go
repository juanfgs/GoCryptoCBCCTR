package ctr

import (
	"testing"
	"encoding/hex"

)

func TestEncrypt(t *testing.T){

	
	for _,n := range EncryptTest {
		Key, _:=hex.DecodeString(n.key)

		c := New([]byte(Key))		
		val := c.Encrypt([]byte(n.input))
		result := c.Decrypt(val)

		if  n.input != string(result) {
			
			t.Fatalf("Expected: %s \n Got: %s", n.input, result)
		}
	}
}

func TestDecrypt(t *testing.T) {
	
	for _,n := range DecryptTest {
		Key, _:=hex.DecodeString(n.key)
		Input, _ :=hex.DecodeString(n.input)		
		c := New([]byte(Key))
		
		if val := string(c.Decrypt(Input)); val != n.expected {
			t.Fatalf("Expected: %s \n Got: %s", n.expected, val)
		}
	}

}

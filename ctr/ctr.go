// Package ctr implements Cipher Block Chaining on Golang
package ctr

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/juanfgs/hmwk_crypto/tools"


)

type CTR struct {
	Cipher cipher.Block
}

func New(key []byte) CTR {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return CTR{cipher}
}

// Encrypt function takes a plain text returns an AES/CBC encrypted CT
func (this CTR) Encrypt(pt []byte ) []byte {
	var iv []byte = make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)

	if err != nil {
		panic(err)
	}
	CT := append([]byte{},iv...)

	max := len(pt) / aes.BlockSize
	rest := len(pt) % aes.BlockSize
	CTChan := make(chan []byte, max +1)
	IVChan := make(chan byte, max +1)
	for i := 0; i <= max; i++ {

		var blockSize int

		if i == max {
			blockSize = rest
		} else {
			blockSize = aes.BlockSize
		}

		block := GrabBlock(pt, blockSize, i * aes.BlockSize)

		go func( Block []byte) {

			CTChan <- tools.XOR(Block, this.EncryptBlock(iv))
			IVChan <- iv[len(iv)-1] + 1

		}( block )

		iv[len(iv)-1] = <- IVChan

	}
	go func(){
		close(CTChan)
		close(IVChan)
	}()


	for val := range CTChan{
		CT = append(CT,val...)
	}

	
	return bytes.TrimSuffix(CT, []byte{0})

}

func (this CTR) Decrypt(ct []byte) []byte {
	var PT []byte
	iv := ct[:aes.BlockSize]

	CT := ct[aes.BlockSize:]


	max := len(CT) / aes.BlockSize
	rest := len(CT) % aes.BlockSize
	PTChan := make(chan []byte, max +1)
	IVChan := make(chan byte, max +1)

	for i := 0; i <= max; i++ {
		var blockSize int

		if i == max {
			blockSize = rest
		} else {
			blockSize = aes.BlockSize
		}

		block := GrabBlock(CT, blockSize, i * aes.BlockSize)
		go func( Block []byte) {

			PTChan <- tools.XOR(Block, this.EncryptBlock(iv))
			IVChan <- iv[len(iv)-1] + 1

		}( block )
		
		iv[len(iv)-1] = <- IVChan

	}
	
	go func(){
		close(PTChan)
		close(IVChan)		
	}()


	for val := range PTChan{
		PT = append(PT,val...)
	}

	return PT
}


func (this CTR) EncryptBlock(ct []byte) []byte {
	var pt []byte = make([]byte, aes.BlockSize)
	this.Cipher.Encrypt(pt, ct)

	return pt
}

// Grabs a block from the full input, specifying  block size and initial index
// Returns the bytes read from the full block
func GrabBlock(full []byte, blockSize int, index int) []byte {

	return full[index : index+blockSize]
}

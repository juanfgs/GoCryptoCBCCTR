// Package ctr implements Cipher Block Chaining on Golang
package ctr

import (
	"crypto/aes"
	"crypto/rand"	
	"crypto/cipher"
	"bytes"
	"github.com/juanfgs/hmwk_crypto/tools"

	"math"
	

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
func (this CTR) Encrypt(pt []byte) []byte {

	var iv []byte = make([]byte, 16)
	_, err := rand.Read(iv)

	if err != nil {
		panic(err)
	}
	CT := iv 

	


	max := len(pt) / aes.BlockSize
	rest := len(pt) % aes.BlockSize
	
	for i:=0; i <= max ; i++ {

		if i == max {
			CT = append(CT,tools.XOR( GrabBlock(pt, rest, i * aes.BlockSize)  , this.chainBlock(iv) )...)
		} else {
			CT = append(CT,tools.XOR( GrabBlock(pt, aes.BlockSize, i * aes.BlockSize)  , this.chainBlock(iv) )...)
		}
		iv[len(iv) -1] = iv[len(iv) -1] + 1		
	}

	return bytes.TrimSuffix(CT, []byte{0})
//	return CT
}


func (this CTR) Decrypt(ct []byte) []byte {

	iv := ct[:aes.BlockSize]

	CT := ct[aes.BlockSize:]

	PT := tools.XOR(this.chainBlock(iv),GrabBlock(CT, aes.BlockSize, 0))
	max := len(CT) / aes.BlockSize
	rest := len(CT) % aes.BlockSize
	for i:=1; i < max; i++ {
		iv[len(iv) -1] = iv[len(iv) -1 ] + 1
		PT = append(PT,tools.XOR( GrabBlock(CT, aes.BlockSize, i * aes.BlockSize )  , this.chainBlock(iv) ) ...)
	}
	iv[len(iv) -1] = iv[len(iv) -1 ] + 1	
	PT = append(PT,tools.XOR( GrabBlock(CT, rest , max * aes.BlockSize )  , this.chainBlock(iv) ) ...)
	return PT
}



func (this CTR) deChainBlock(ct []byte) []byte {
	var pt []byte = make([]byte, aes.BlockSize)
	this.Cipher.Decrypt(pt, ct)

	return pt
}

func (this CTR) chainBlock(src []byte) []byte {
	var ct []byte = make([]byte, aes.BlockSize)

	this.Cipher.Encrypt(ct, src)
	return ct
}

// Pads to fit blocksize

func Pad(origin []byte, blockSize int) []byte {
	bitblocks := int(math.Ceil(float64(len(origin)) / float64(16)))

	var dst []byte = make([]byte, bitblocks*blockSize)
	copy(dst, origin)

	return dst
}

// Grabs a block from the full input, specifying  block size and initial index
// Returns the bytes read from the full block
func GrabBlock(full []byte, blockSize int, index int) []byte {

	return full[index : index+blockSize]
}

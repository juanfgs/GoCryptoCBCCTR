// Package cbc implements Cipher Block Chaining on Golang
package cbc

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/juanfgs/hmwk_crypto/tools"
	"math"
)

type CBC struct {
	Cipher cipher.Block
}

func New(key []byte) CBC {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return CBC{cipher}
}

// Encrypt function takes a plain text returns an AES/CBC encrypted CT
func (this CBC) Encrypt(pt []byte) []byte {
	/* REIMPLEMENT THIS */
	return pt
}

func (this CBC) Decrypt(ct []byte) []byte {
	iv := ct[:aes.BlockSize]
	CT := ct[aes.BlockSize:]

	PT := tools.XOR(this.deChainBlock(GrabBlock(CT, aes.BlockSize, 0)) ,iv)
	max := len(CT) / aes.BlockSize

	for i:=1; i < max; i++ {
		prev := GrabBlock(CT, aes.BlockSize, (i -1) * aes.BlockSize )
		current := this.deChainBlock(GrabBlock(CT, aes.BlockSize, i * aes.BlockSize ))

		PT = append(PT,tools.XOR(current,prev)...)
	}
	pad := PT[len(PT)-1:]

	return PT[:byte(len(PT))-pad[0]]
}



func (this CBC) deChainBlock(ct []byte) []byte {
	var pt []byte = make([]byte, aes.BlockSize)
	this.Cipher.Decrypt(pt, ct)

	return pt
}

func (this CBC) chainBlock(src []byte) []byte {
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

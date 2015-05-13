package cbc

var PadTest = []struct{
	input []byte
	blockSize int
	expect int
} {
	{[]byte("foobar"), 16, 0},
	{[]byte("fdasfdsafdsafdsafdsaffofdasasobar"), 32, 0},	
}

var DecryptTest = []struct {
	cipher CBC
	input []byte
	expected string
	
} {
	{ New([]byte("140b41b22a29beb4061bda66b6747e14")) , []byte{192,88,120,204,71,185,195,167,188,48,22,196,62,141,88,31,239,216,252,145,171,64,163,250,230,81,158,218,195,89,187,62},"foobar"  },

	
}


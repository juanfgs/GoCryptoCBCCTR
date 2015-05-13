package ctr


var DecryptTest = []struct {
	key string
	input string
	expected string
	
} {
	{ "36f18357be4dbd77f050515c73fcf9f2" , "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329","CTR mode lets you build a stream cipher from a block cipher."  },
	{ "36f18357be4dbd77f050515c73fcf9f2","770b80259ec33beb2561358a9f2dc617e46218c0a53cbeca695ae45faa8952aa0e311bde9d4e01726d3184c34451","Always avoid the two time pad!"},
}


var EncryptTest = []struct {
	key string
	input string

	
} {
	{ "36f18357be4dbd77f050515c73fcf9f2", "CTR mode lets you build a stream cipher from a block cipher."},
	{ "36f18357be4dbd77f050515c73fcf9f2", "Always avoid the two time pad!"},
	{ "36f18357be4dbd77f050515c73fcf9f2", "Testing various types of variable messages     "},
	
}

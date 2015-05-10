package cryptools

var XORTests = []struct {
	a, b string
	expected string
} {
	{"deadbeef","beefdead","\x06\x00\x04\x02\x06\x00\x04\x02"},
	{"vamosamedirnos","averquienesmas", "\x17\x17\b\x1D\x02\x14\x04\x00\n\f\x01\x03\x0E\x00"},	
}


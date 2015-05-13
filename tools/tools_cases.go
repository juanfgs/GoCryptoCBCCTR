package tools

var XORTests = []struct {
	a, b string
	expected string
} {
	{"deadbeef","beefdead","\x06\x00\x04\x02\x06\x00\x04\x02"},
	{"deadbeef","deadbeef","\x00\x00\x00\x00\x00\x00\x00\x00"},	
	{"d","0", "d"},

}


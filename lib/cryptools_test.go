package cryptools

import(
	"testing"
)

const testVersion = 2

func TestXOR(t *testing.T) {
	for _, n := range XORTests {
		if result := XOR(n.a,n.b); result != n.expected {
			t.Fatalf("Expected %s, got %s", n.expected, result)
		}
	}
}


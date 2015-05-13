package tools

import(
	"testing"
)

const testVersion = 2

func TestXOR(t *testing.T) {
	for _, n := range XORTests {
		if result := XOR([]byte(n.a),[]byte(n.b)); string(result) != n.expected {
			t.Fatalf("Expected %s, got %s", n.expected, result)
		}
	}
}


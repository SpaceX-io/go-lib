package rand

import "testing"

func TestString(t *testing.T) {
	s := String(20)
	t.Logf("rand string: %s", s)
}

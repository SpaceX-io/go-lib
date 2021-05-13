package snowflake

import "testing"

func TestInt64(t *testing.T) {
	node, err := NewNode(1)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	id := node.Generate()

	t.Logf("Int64 : %#v", id.Int64())
}

func TestString(t *testing.T) {
	node, err := NewNode(1)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	id := node.Generate()

	t.Logf("String : %#v", id.String())
}

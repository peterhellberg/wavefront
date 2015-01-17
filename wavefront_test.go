package wavefront

import "testing"

func TestRead(t *testing.T) {
	model, err := Read("models/gopher.obj")
	if err != nil {
		t.Fatalf("%v", err)
	}

	if len(model) != 12 {
		t.Fatalf("Expected number of objects to be %d", 12)
	}
}

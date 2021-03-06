package storagetesting

import "testing"

func TestMapStorage(t *testing.T) {
	data := map[string][]byte{}
	r := NewMapStorage(data)
	if r == nil {
		t.Errorf("unexpected result: %v", r)
	}
	VerifyStorage(t, r)
}

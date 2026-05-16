package practice

import "testing"

func TestEx002(t *testing.T) {
	if _, err := Ex002(-10); err == nil {
		t.Errorf("Ex002(-10) expected error, got nil")
	}

	if got, err := Ex002(0); err != nil || got != 1 {
		t.Errorf("Ex002(0) = (%v, %v), want (1, nil)", got, err)
	}

	if got, err := Ex002(8); err != nil || got != 40320 {
		t.Errorf("Ex002(8) = (%v, %v), want (40320, nil)", got, err)
	}
}

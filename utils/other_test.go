package utils

import "testing"

// test StructSpaceTrim
func TestStructSpaceTrim(t *testing.T) {
	// test case 1
	type TestStruct struct {
		Name string
		Age  int
	}
	ts := TestStruct{
		Name: "  test  ",
	}
	StructSpaceTrim(&ts)
	if ts.Name != "test" {
		t.Errorf("test case 1 failed")
	}
}

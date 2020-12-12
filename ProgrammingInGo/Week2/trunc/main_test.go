package main

import "testing"

func Test_convertFloatToInt(t *testing.T) {
	tests := []struct {
		name          string
		inputFloatStr string
		expectedInt   int
	}{
		{name: "simple float", inputFloatStr: "15.3", expectedInt: 15},
		{name: "input string", inputFloatStr: "test", expectedInt: -1},
		{name: "input int", inputFloatStr: "7676", expectedInt: 7676},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertFloatToInt(tt.inputFloatStr)
			if got != tt.expectedInt {
				t.Errorf("Expected %d; Got %d\n", tt.expectedInt, got)
			} else {
				t.Logf("Expected %d; Got %d\n", tt.expectedInt, got)
			}
		})
	}
}

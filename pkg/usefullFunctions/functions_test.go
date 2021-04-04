package usefullFunctions

import "testing"

var needToResponse = []int32{2, 3, 5, 8, 13, 21, 34, 55}

type testvalidate struct {
	values []int32
	resp   bool
}

func TestValidate(t *testing.T) {
	testValues := []testvalidate{
		{
			values: []int32{1, 2},
			resp:   true,
		},
		{
			values: []int32{-1, 10},
			resp:   false,
		},
		{
			values: []int32{10, 9},
			resp:   false,
		},
		{
			values: []int32{3, -4},
			resp:   false,
		},
		{
			values: []int32{1, 100},
			resp:   true,
		},
		{
			values: []int32{-5, -1},
			resp:   false,
		},
	}
	for _, value := range testValues {
		if value.resp != Validate(value.values[0], value.values[1]) {
			t.Errorf("Something went wrong")
		}
	}
}

func TestCalculate(t *testing.T) {
	x1, x2 := 1, 1
	count := int32(8)
	resp := Calculate(x1, x2, count)
	for i, _ := range resp {
		if resp[i] != needToResponse[i] {
			t.Error("Returned response:", resp, " . What i wanna see: ", needToResponse)
		}
	}
}

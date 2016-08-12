package problem

import (
	"strconv"
	"testing"
)

func TestIsPrime(t *testing.T) {
	var num1 int64 = -1
	assert1 := false
	if assert1 != IsPrime(num1) {
		t.Errorf(strconv.FormatInt(num1, 10) + " Testing fails")
	}

	var num2 int64 = 0
	assert2 := false
	if assert2 != IsPrime(num2) {
		t.Errorf(strconv.FormatInt(num2, 10) + " Testing fails")
	}

	var num3 int64 = 1
	assert3 := false
	if assert3 != IsPrime(num3) {
		t.Errorf(strconv.FormatInt(num3, 10) + " Testing fails")
	}

	var num4 int64 = 2
	assert4 := true
	if assert4 != IsPrime(num4) {
		t.Errorf(strconv.FormatInt(num4, 10) + " Testing fails")
	}

	var num5 int64 = 3
	assert5 := true
	if assert5 != IsPrime(num5) {
		t.Errorf(strconv.FormatInt(num5, 10) + " Testing fails")
	}

	var num6 int64 = 16
	assert6 := false
	if assert6 != IsPrime(num6) {
		t.Errorf(strconv.FormatInt(num6, 10) + " Testing fails")
	}

	var num7 int64 = 17
	assert7 := true
	if assert7 != IsPrime(num7) {
		t.Errorf(strconv.FormatInt(num7, 10) + " Testing fails")
	}

}

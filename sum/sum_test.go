package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{1, 2, 3},
	{4, 5, 9},
	{126, 1, 127},
	{-127, -1, -127},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_uint32 = []struct {
	n1		 uint32
	n2		 uint32
	expected uint32
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{4294967294, 1, 4294967294},
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
			
		}
	}
}

var sum_tests_int64 = []struct {
	n1		 int64
	n2		 int64
	expected int64
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{9223372036854775806, 1, 9223372036854775806},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
			
		}
	}
}

var sum_tests_int32 = []struct {
	n1		 int32
	n2		 int32
	expected int32
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{2147483646, 1, 2147483646},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
			
		}
	}
}

var sum_tests_float64 = []struct {
	n1		 float64
	n2		 float64
	expected float64
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
	{1.000785, 1, 2},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%f, %f) returned %f, expected %f", v.n1, v.n2, val, v.expected)
			
		}
	}
}
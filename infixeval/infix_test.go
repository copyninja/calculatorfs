package infix

import (
	// "fmt"
	"testing"
)

type testData struct {
	expression   string
	result_type  string
	value_int    int64
	value_float  float64
	error_string string
}

func (testcase *testData) evaluateInfixExpr(t *testing.T) {
	var value interface{}
	var err error
	if value, err = Evaluate(testcase.expression); err != nil {
		if err.Error() != testcase.error_string {
			t.Errorf("Was expecting error: %s, but got %s", testcase.error_string, err)
		}
	}

	switch testcase.result_type {
	case "int64":
		result, ok := value.(int64)
		if !ok {
			t.Errorf("Was expecting int64 result but failed to type assert")
		}

		if result != testcase.value_int {
			t.Errorf("'%s' = %d but got %d", testcase.expression, testcase.value_int, result)
		}
	case "float64":
		result, ok := value.(float64)
		if !ok {
			t.Errorf("Was expecting float64 value but failed to type assert")
		}

		if result != testcase.value_float {
			t.Errorf("'%s' = %f but got %f", testcase.expression, testcase.value_float, result)
		}

	}
}

func TestSimpleEval(t *testing.T) {
	input := []*testData{
		{"2 + 3", "int64", 5, 0, ""},
		{"2.0 + 3", "float64", 0, 5.0, ""},
		{"2", "int64", 2, 0, ""},
		{"2 - 3", "int64", -1, 0, ""},
		{"2 * 3", "int64", 6, 0, ""},
		{"2 / 3", "int64", 0, 0, ""},
		{"2 / 4.0", "float64", 0, 0.5, ""},
		{"5 * 4.0", "float64", 0, 20.0, ""},
		{"20.0 / 4.0", "float64", 0, 5.0, ""},
		{"4.0 - 5.0", "float64", 0, -1.0, ""},

		// Simple Errors
		{"2 / 0", "eror", 0, 0, "Attempt to divide by 0"},
		{"2 % 0", "error", 0, 0, "Attempt to divide by 0"},
		{"8.0 % 2", "error", 0, 0, "Modulus operator is not supported on float"},
		{"( 2 + 3", "error", 0, 0, "Invalid/unbalanced expression"},
		{"(2 + hello)", "error", 0, 0, "Invalid tokens in the expression"},

		// Nested errors
		{"(2 / 0)", "error", 0, 0, "Attempt to divide by 0"},
		{"2 % 0 + 3", "error", 0, 0, "Attempt to divide by 0"},
	}

	for _, test := range input {
		test.evaluateInfixExpr(t)
	}
}

func TestComplexEval(t *testing.T) {
	input := []*testData{
		{"2 + (3*5) % 10", "int64", 7, 0, ""},
		{"(4 * 15) % 10 / 2", "int64", 0, 0, ""},
		{"(2*3 + (5*8) % 28)", "int64", 18, 0, ""},

		{"11 * 6 + 21 )", "error", 0, 0, "Invalid/unbalanced expression"},
	}

	for _, test := range input {
		test.evaluateInfixExpr(t)
	}
}

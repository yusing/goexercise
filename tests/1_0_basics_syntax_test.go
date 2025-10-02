package goexercises

import (
	"goexercises"
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		name  string
		x, y  int
		wantX int
		wantY int
	}{
		{
			name:  "positive numbers",
			x:     1,
			y:     2,
			wantX: 2,
			wantY: 1,
		},
		{
			name:  "negative numbers",
			x:     -1,
			y:     -2,
			wantX: -2,
			wantY: -1,
		},
		{
			name:  "zero values",
			x:     0,
			y:     5,
			wantX: 5,
			wantY: 0,
		},
		{
			name:  "same values",
			x:     3,
			y:     3,
			wantX: 3,
			wantY: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := goexercises.Swap(tt.x, tt.y)
			if gotX != tt.wantX || gotY != tt.wantY {
				t.Errorf("Swap(%d, %d) = (%d, %d), want (%d, %d)",
					tt.x, tt.y, gotX, gotY, tt.wantX, tt.wantY)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected string
	}{
		{
			name:     "simple strings",
			a:        "hello",
			b:        "world",
			expected: "helloworld",
		},
		{
			name:     "empty strings",
			a:        "",
			b:        "",
			expected: "",
		},
		{
			name:     "one empty string",
			a:        "hello",
			b:        "",
			expected: "hello",
		},
		{
			name:     "with spaces",
			a:        "hello ",
			b:        "world",
			expected: "hello world",
		},
		{
			name:     "special characters",
			a:        "foo",
			b:        "bar!",
			expected: "foobar!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.Concat(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Concat(%q, %q) = %q, want %q",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestConcatN(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "no arguments",
			input:    []string{},
			expected: "",
		},
		{
			name:     "single string",
			input:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "two strings",
			input:    []string{"hello", "world"},
			expected: "helloworld",
		},
		{
			name:     "multiple strings",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: "abcde",
		},
		{
			name:     "with spaces",
			input:    []string{"hello", " ", "world"},
			expected: "hello world",
		},
		{
			name:     "empty strings mixed",
			input:    []string{"", "hello", "", "world", ""},
			expected: "helloworld",
		},
		{
			name:     "special characters",
			input:    []string{"foo", "!", "@", "#"},
			expected: "foo!@#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.ConcatN(tt.input...)
			if result != tt.expected {
				t.Errorf("ConcatN(%v) = %q, want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		s, sep   string
		expected []string
	}{
		{
			name:     "simple comma separated",
			s:        "1,2,3",
			sep:      ",",
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "space separated",
			s:        "hello world go",
			sep:      " ",
			expected: []string{"hello", "world", "go"},
		},
		{
			name:     "no separator",
			s:        "hello",
			sep:      ",",
			expected: []string{"hello"},
		},
		{
			name:     "empty string",
			s:        "",
			sep:      ",",
			expected: []string{""},
		},
		{
			name:     "multiple separators",
			s:        "a,,b,,c",
			sep:      ",",
			expected: []string{"a", "", "b", "", "c"},
		},
		{
			name:     "separator at start",
			s:        ",hello,world",
			sep:      ",",
			expected: []string{"", "hello", "world"},
		},
		{
			name:     "separator at end",
			s:        "hello,world,",
			sep:      ",",
			expected: []string{"hello", "world", ""},
		},
		{
			name:     "multi-character separator",
			s:        "apple::banana::cherry",
			sep:      "::",
			expected: []string{"apple", "banana", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.Split(tt.s, tt.sep)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Split(%q, %q) = %v, want %v",
					tt.s, tt.sep, result, tt.expected)
			}
		})
	}
}

func TestTypeAssertion1(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected int
	}{
		{
			name:     "string value",
			input:    "hello",
			expected: 1,
		},
		{
			name:     "int value",
			input:    42,
			expected: 2,
		},
		{
			name:     "float64 value",
			input:    3.14,
			expected: 3,
		},
		{
			name:     "bool value",
			input:    true,
			expected: 0,
		},
		{
			name:     "slice value",
			input:    []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "map value",
			input:    map[string]int{"a": 1},
			expected: 0,
		},
		{
			name:     "nil value",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 1,
		},
		{
			name:     "zero int",
			input:    0,
			expected: 2,
		},
		{
			name:     "zero float64",
			input:    0.0,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.TypeAssertion1(tt.input)
			if result != tt.expected {
				t.Errorf("TypeAssertion1(%v) = %d, want %d",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestTypeAssertion2(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		expectedStr string
		expectedOk  bool
	}{
		{
			name:        "string value",
			input:       "hello world",
			expectedStr: "hello world",
			expectedOk:  true,
		},
		{
			name:        "int value",
			input:       42,
			expectedStr: "",
			expectedOk:  false,
		},
		{
			name:        "float64 value",
			input:       3.14,
			expectedStr: "",
			expectedOk:  false,
		},
		{
			name:        "bool value",
			input:       true,
			expectedStr: "",
			expectedOk:  false,
		},
		{
			name:        "slice value",
			input:       []int{1, 2, 3},
			expectedStr: "",
			expectedOk:  false,
		},
		{
			name:        "map value",
			input:       map[string]int{"a": 1},
			expectedStr: "",
			expectedOk:  false,
		},
		{
			name:        "nil value",
			input:       nil,
			expectedStr: "",
			expectedOk:  false,
		},
		{
			name:        "empty string",
			input:       "",
			expectedStr: "",
			expectedOk:  true,
		},
		{
			name:        "string with special characters",
			input:       "hello@world.com",
			expectedStr: "hello@world.com",
			expectedOk:  true,
		},
		{
			name:        "string with spaces",
			input:       "hello world",
			expectedStr: "hello world",
			expectedOk:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultStr, resultOk := goexercises.TypeAssertion2(tt.input)
			if resultStr != tt.expectedStr || resultOk != tt.expectedOk {
				t.Errorf("TypeAssertion2(%v) = (%q, %v), want (%q, %v)",
					tt.input, resultStr, resultOk, tt.expectedStr, tt.expectedOk)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "positive number",
			input:    42,
			expected: "42",
		},
		{
			name:     "negative number",
			input:    -123,
			expected: "-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.IntToString(tt.input)
			if result != tt.expected {
				t.Errorf("IntToString(%d) = %q, want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    int
		expectedErr bool
	}{
		{
			name:        "positive number",
			input:       "42",
			expected:    42,
			expectedErr: false,
		},
		{
			name:        "negative number",
			input:       "-123",
			expected:    -123,
			expectedErr: false,
		},
		{
			name:        "zero",
			input:       "0",
			expected:    0,
			expectedErr: false,
		},
		{
			name:        "invalid number",
			input:       "invalid",
			expected:    0,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := goexercises.StringToInt(tt.input)
			if tt.expectedErr {
				if err == nil {
					t.Errorf("StringToInt(%q) = %d, nil, want error",
						tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt(%q) = %d, %v, want %d, nil",
						tt.input, result, err, tt.expected)
				}
				if result != tt.expected {
					t.Errorf("StringToInt(%q) = %d, want %d",
						tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestSliceCopy(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "simple slice",
			input: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.SliceCopy(tt.input)
			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("SliceCopy(%v) = %v, want %v", tt.input, result, tt.input)
			}
			if &result[0] == &tt.input[0] {
				t.Errorf("SliceCopy(%v) = %v should return a new slice", tt.input, result)
			}
		})
	}
}

func TestMapCopy(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
	}{
		{
			name:  "simple map",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.MapCopy(tt.input)
			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("MapCopy(%v) = %v, want %v", tt.input, result, tt.input)
			}
			if reflect.ValueOf(result).Pointer() == reflect.ValueOf(tt.input).Pointer() {
				t.Errorf("MapCopy(%v) = %v should return a new map", tt.input, result)
			}
		})
	}
}

func TestFormatString(t *testing.T) {
	tests := []struct {
		name     string
		time     string
		title    string
		msg      string
		expected string
	}{
		{
			name:     "simple string",
			time:     "2021-01-01 12:00:00",
			title:    "Hello",
			msg:      "World",
			expected: "[2021-01-01 12:00:00] Hello: World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goexercises.FormatString(tt.time, tt.title, tt.msg)
			if result != tt.expected {
				t.Errorf("FormatString(%q, %q, %q) = %q, want %q", tt.time, tt.title, tt.msg, result, tt.expected)
			}
		})
	}
}

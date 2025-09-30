package goexercises

import (
	"errors"
	"goexercises"
	"strings"
	"testing"
	"time"
)

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative result", -10, 2, -5, false},
		{"float result", 7, 2, 3.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := goexercises.SafeDivide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeDivide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("SafeDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindFirst(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		predicate func(int) bool
		want      int
		wantOk    bool
	}{
		{
			name:      "find even number",
			items:     []int{1, 3, 4, 5},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      4,
			wantOk:    true,
		},
		{
			name:      "not found",
			items:     []int{1, 3, 5},
			predicate: func(n int) bool { return n%2 == 0 },
			want:      0,
			wantOk:    false,
		},
		{
			name:      "empty slice",
			items:     []int{},
			predicate: func(n int) bool { return true },
			want:      0,
			wantOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := goexercises.FindFirst(tt.items, tt.predicate)
			if ok != tt.wantOk {
				t.Errorf("FindFirst() ok = %v, want %v", ok, tt.wantOk)
				return
			}
			if got != tt.want {
				t.Errorf("FindFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	goexercises.MapValues(m, func(n int) int { return n * 2 })

	want := map[string]int{"a": 2, "b": 4, "c": 6}
	for k, v := range want {
		if m[k] != v {
			t.Errorf("MapValues() key %s = %v, want %v", k, m[k], v)
		}
	}
}

func TestDebounce(t *testing.T) {
	var result string
	fn := func(s string) {
		result = s
	}

	debounced := goexercises.Debounce(fn, 50)

	debounced("first")
	debounced("second")
	debounced("third")

	time.Sleep(30 * time.Millisecond)
	if result != "" {
		t.Errorf("Debounce() executed too early, got %v", result)
	}

	time.Sleep(30 * time.Millisecond)
	if result != "third" {
		t.Errorf("Debounce() = %v, want 'third'", result)
	}
}

func TestParseKeyValue(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    map[string]string
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "name=John;age=30;city=NYC",
			want:  map[string]string{"name": "John", "age": "30", "city": "NYC"},
		},
		{
			name:  "empty string",
			input: "",
			want:  map[string]string{},
		},
		{
			name:    "invalid format - no equals",
			input:   "name:John",
			wantErr: true,
		},
		{
			name:    "invalid format - multiple equals",
			input:   "name=John=Doe",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := goexercises.ParseKeyValue(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseKeyValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("ParseKeyValue() length = %v, want %v", len(got), len(tt.want))
					return
				}
				for k, v := range tt.want {
					if got[k] != v {
						t.Errorf("ParseKeyValue() key %s = %v, want %v", k, got[k], v)
					}
				}
			}
		})
	}
}

func TestRetry(t *testing.T) {
	t.Run("succeeds on first attempt", func(t *testing.T) {
		attempts := 0
		fn := func() error {
			attempts++
			return nil
		}

		err := goexercises.Retry(fn, 3)
		if err != nil {
			t.Errorf("Retry() error = %v, want nil", err)
		}
		if attempts != 1 {
			t.Errorf("Retry() attempts = %v, want 1", attempts)
		}
	})

	t.Run("succeeds on second attempt", func(t *testing.T) {
		attempts := 0
		fn := func() error {
			attempts++
			if attempts < 2 {
				return errors.New("temporary error")
			}
			return nil
		}

		err := goexercises.Retry(fn, 3)
		if err != nil {
			t.Errorf("Retry() error = %v, want nil", err)
		}
		if attempts != 2 {
			t.Errorf("Retry() attempts = %v, want 2", attempts)
		}
	})

	t.Run("fails after max attempts", func(t *testing.T) {
		attempts := 0
		testErr := errors.New("permanent error")
		fn := func() error {
			attempts++
			return testErr
		}

		err := goexercises.Retry(fn, 3)
		if err == nil {
			t.Error("Retry() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "permanent error") {
			t.Errorf("Retry() error = %v, want error containing 'permanent error'", err)
		}
		if attempts != 3 {
			t.Errorf("Retry() attempts = %v, want 3", attempts)
		}
	})
}

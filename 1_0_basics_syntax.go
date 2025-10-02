// Basics Syntax and standard library usage
package goexercises

// Swap swaps two strings and returns them in reverse order
func Swap(x, y int) (int, int) {
	// TODO: implement
	return 0, 0
}

func Concat(a, b string) string {
	// TODO: implement
	return ""
}

func ConcatN(strings ...string) string {
	// TODO: implement
	// tips:
	// `...strings` is a variadic parameter, it's type is []string
	// Caller will pass in a variable number of strings, e.g. ConcatN("a", "b", "c"), ConcatN("a", "b", "c", "d", "e").
	//
	// Use strings.Builder for efficient string concatenation.
	// Just type the module name (e.g. strings.) to autocomplete,
	// enable format on save to automatically import the module on file save.
	return ""
}

func Split(s, sep string) []string {
	// TODO: implement
	// Split("1,2,3", ",") should return []string{"1", "2", "3"}
	return []string{}
}

func TypeAssertion1(v any) int {
	// TODO: implement
	/**
	Return:
		- 1 if the value is a string
		- 2 if the value is a int
		- 3 if the value is a float64
		- 0 if the value is not a string, int, or float64
	*/
	return 0
}

func TypeAssertion2(v any) (string, bool) {
	// TODO: implement
	/**
	Return:
		- string value and true if the value is a string
		- otherwise return empty string and false
	*/
	return "", false
}

func IntToString(i int) string {
	// TODO: implement
	// return the string representation of the integer
	return ""
}

func StringToInt(s string) (int, error) {
	// TODO: implement
	// return the integer representation of the string
	return 0, nil
}

func SliceCopy(src []int) []int {
	// TODO: implement
	// return a copy of the slice, underlying pointer should be different
	return []int{}
}

func MapCopy(src map[string]int) map[string]int {
	// TODO: implement
	// return a copy of the map, underlying pointer should be different
	return map[string]int{}
}

func FormatString(time, title, msg string) string {
	// TODO: implement
	// return the string in format of "[time] title: msg"
	//
	// tips: use fmt.Sprintf to format the string
	return ""
}

package goexercises

// Cache is a generic-like cache using any type.
// Teaches working with interface{} and type assertions.
type Cache struct {
	data map[string]any
}

// NewCache creates a new cache instance.
func NewCache() *Cache {
	// TODO: implement
	return nil
}

// Set stores a value in the cache.
func (c *Cache) Set(key string, value any) {
	// TODO: implement
}

// Get retrieves a value from cache and type-asserts it.
// Returns the value and true if found and type matches, zero value and false otherwise.
func (c *Cache) Get(key string) (any, bool) {
	// TODO: implement
	return nil, false
}

// GetString is a convenience method that gets and type-asserts to string.
func (c *Cache) GetString(key string) (string, bool) {
	// TODO: implement
	return "", false
}

// Stack is a generic stack implementation.
// Teaches Go generics (new feature frontend devs with TS knowledge will appreciate).
type Stack[T any] struct {
	items []T
}

// NewStack creates a new generic stack.
func NewStack[T any]() *Stack[T] {
	// TODO: implement
	return nil
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
	// TODO: implement
}

// Pop removes and returns the top item.
// Returns zero value and false if stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	// TODO: implement
	return zero, false
}

// IsEmpty returns true if stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	// TODO: implement
	return false
}

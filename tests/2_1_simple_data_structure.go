package goexercises

import (
	. "goexercises"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	cache.Set("name", "John")
	cache.Set("age", 25)
	cache.Set("active", true)

	t.Run("Get existing values", func(t *testing.T) {
		val, ok := cache.Get("name")
		if !ok {
			t.Error("Get() should find 'name'")
		}
		if val != "John" {
			t.Errorf("Get() = %v, want 'John'", val)
		}
	})

	t.Run("Get non-existing key", func(t *testing.T) {
		_, ok := cache.Get("nonexistent")
		if ok {
			t.Error("Get() should not find non-existent key")
		}
	})

	t.Run("GetString", func(t *testing.T) {
		str, ok := cache.GetString("name")
		if !ok {
			t.Error("GetString() should find 'name'")
		}
		if str != "John" {
			t.Errorf("GetString() = %v, want 'John'", str)
		}

		_, ok = cache.GetString("age")
		if ok {
			t.Error("GetString() should return false for non-string value")
		}
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stack := NewStack[int]()

		if !stack.IsEmpty() {
			t.Error("New stack should be empty")
		}

		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		if stack.IsEmpty() {
			t.Error("Stack with items should not be empty")
		}

		val, ok := stack.Pop()
		if !ok || val != 3 {
			t.Errorf("Pop() = (%v, %v), want (3, true)", val, ok)
		}

		val, ok = stack.Pop()
		if !ok || val != 2 {
			t.Errorf("Pop() = (%v, %v), want (2, true)", val, ok)
		}
	})

	t.Run("string stack", func(t *testing.T) {
		stack := NewStack[string]()

		stack.Push("hello")
		stack.Push("world")

		val, ok := stack.Pop()
		if !ok || val != "world" {
			t.Errorf("Pop() = (%v, %v), want (world, true)", val, ok)
		}
	})

	t.Run("pop from empty stack", func(t *testing.T) {
		stack := NewStack[int]()
		_, ok := stack.Pop()
		if ok {
			t.Error("Pop() from empty stack should return false")
		}
	})
}

func TestStackStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	stack := NewStack[Person]()
	stack.Push(Person{"Alice", 30})
	stack.Push(Person{"Bob", 25})

	person, ok := stack.Pop()
	if !ok || person.Name != "Bob" {
		t.Errorf("Pop() = %v, want Bob", person)
	}
}

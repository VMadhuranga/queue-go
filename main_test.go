package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	cases := []struct {
		operations []struct {
			method string
			data   any
		}
		expected []any
	}{
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", "Rand"},
				{"push", "Mat"},
				{"peek", nil},
				{"pop", nil},
			},
			expected: []any{
				nil,
				nil,
				"Rand",
				"Rand",
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", "Egwene"},
				{"push", "Nynaeve"},
				{"size", nil},
				{"pop", nil},
				{"size", nil},
			},
			expected: []any{
				nil,
				nil,
				2,
				"Egwene",
				1,
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{

				{"pop", nil},
				{"peek", nil},
				{"size", nil},
			},
			expected: []any{
				nil,
				nil,
				0,
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{

				{"push", "Perrin"},
				{"push", "Moiraine"},
				{"push", "Lan"},
				{"pop", nil},
				{"pop", nil},
				{"peek", nil},
			},
			expected: []any{
				nil,
				nil,
				nil,
				"Perrin",
				"Moiraine",
				"Lan",
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"push", "Thom"},
				{"pop", nil},
				{"push", "Loial"},
				{"peek", nil},
			},
			expected: []any{
				nil,
				"Thom",
				nil,
				"Loial",
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"pop", nil},
				{"peek", nil},
				{"size", nil},
				{"push", "Thom"},
				{"size", nil},
				{"peek", nil},
				{"pop", nil},
				{"size", nil},
				{"push", "Thom"},
				{"size", nil},
				{"peek", nil},
				{"pop", nil},
				{"size", nil},
			},
			expected: []any{
				nil,
				nil,
				0,
				nil,
				1,
				"Thom",
				"Thom",
				0,
				nil,
				1,
				"Thom",
				"Thom",
				0,
			},
		},
	}

	for _, c := range cases {
		q := queue{}

		for i, o := range c.operations {
			switch o.method {
			case "push":
				q.push(o.data)
			case "pop":
				got, expected := q.pop(), c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			case "peek":
				got, expected := q.peek(), c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			case "size":
				got, expected := q.getSize(), c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			}
		}
	}
}

func testHelper(t *testing.T, idx int, got, expected any) {
	t.Errorf("expected IDX: %v | got: %v != expected: %v", idx, got, expected)
}

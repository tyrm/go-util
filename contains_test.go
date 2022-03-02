package util

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestContainsString(t *testing.T) {
	stack := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	tables := []struct {
		x string
		n bool
	}{
		{"one", true},
		{"four", true},
		{"foo", false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Checking stack for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ContainsString(stack, table.x)
			if result != table.n {
				t.Errorf("Sum was incorrect, got: %v, want: %v.", result, table.n)
			}
		})
	}
}

func TestContainsOneOfStrings(t *testing.T) {
	stack := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	tables := []struct {
		x []string
		n bool
	}{
		{[]string{"one", "two", "three"}, true},
		{[]string{"foo", "five", "bar"}, true},
		{[]string{"foo", "bar", "fizz"}, false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Checking stack for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ContainsOneOfStrings(stack, table.x)
			if result != table.n {
				t.Errorf("Sum was incorrect, got: %v, want: %v.", result, table.n)
			}
		})
	}
}

func TestContainsUUID(t *testing.T) {
	stack := []uuid.UUID{
		uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138")),
		uuid.Must(uuid.Parse("e420e8b4-3873-43bc-a3d4-b5b0211754b9")),
		uuid.Must(uuid.Parse("f9ef99a3-cb13-4688-8547-c78081053dca")),
		uuid.Must(uuid.Parse("969b3d8f-a03d-4016-8202-d57ea8eae49f")),
		uuid.Must(uuid.Parse("92044b18-91fd-4689-861d-99ea543d4191")),
	}

	tables := []struct {
		x uuid.UUID
		n bool
	}{
		{uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138")), true},
		{uuid.Must(uuid.Parse("969b3d8f-a03d-4016-8202-d57ea8eae49f")), true},
		{uuid.Must(uuid.Parse("319df288-032e-4628-ac4d-be483f263c37")), false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Checking stack for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ContainsUUID(stack, table.x)
			if result != table.n {
				t.Errorf("Sum was incorrect, got: %v, want: %v.", result, table.n)
			}
		})
	}
}

func TestContainsOneOfUUIDs(t *testing.T) {
	stack := []uuid.UUID{
		uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138")),
		uuid.Must(uuid.Parse("e420e8b4-3873-43bc-a3d4-b5b0211754b9")),
		uuid.Must(uuid.Parse("f9ef99a3-cb13-4688-8547-c78081053dca")),
		uuid.Must(uuid.Parse("969b3d8f-a03d-4016-8202-d57ea8eae49f")),
		uuid.Must(uuid.Parse("92044b18-91fd-4689-861d-99ea543d4191")),
	}

	tables := []struct {
		x []uuid.UUID
		n bool
	}{
		{
			[]uuid.UUID{
				uuid.Must(uuid.Parse("9be91ef7-e1b9-46e0-9418-44f5e5d5b138")),
				uuid.Must(uuid.Parse("e420e8b4-3873-43bc-a3d4-b5b0211754b9")),
				uuid.Must(uuid.Parse("f9ef99a3-cb13-4688-8547-c78081053dca")),
			},
			true,
		},
		{
			[]uuid.UUID{
				uuid.Must(uuid.Parse("e0f30d4a-250b-4200-9d0f-b057a820e58b")),
				uuid.Must(uuid.Parse("92044b18-91fd-4689-861d-99ea543d4191")),
				uuid.Must(uuid.Parse("ac0c8444-eb3b-4a6c-97dc-219f07e66c6c")),
			},
			true,
		},
		{
			[]uuid.UUID{
				uuid.Must(uuid.Parse("2d43417a-5da6-4123-80d6-2f2d42f7477a")),
				uuid.Must(uuid.Parse("4a90659a-423d-471e-baa1-b87a0a66c51a")),
				uuid.Must(uuid.Parse("03c6f272-bb38-486f-bde1-eba3ac6d6ff6")),
			},
			false,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Checking stack for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ContainsOneOfUUIDs(stack, table.x)
			if result != table.n {
				t.Errorf("Sum was incorrect, got: %v, want: %v.", result, table.n)
			}
		})
	}
}

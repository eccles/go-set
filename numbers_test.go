package set_test

import (
	"fmt"

	"github.com/eccles/go-set"
)

func ExampleTruncateList_midduplicates() {
	a := []int{10, 20, 20, 20, 30}
	b := set.TruncateList(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 2 [10 30]
}

func ExampleTruncateList_midduplicates1() {
	a := []int{10, 20, 20, 30}
	b := set.TruncateList(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 4 [10 20 20 30]
}

func ExampleTruncateList_trailingduplicate() {
	a := []int{10, 20, 30, 30, 30}
	b := set.TruncateList(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 2 [10 20]
}

func ExampleTruncateList_trailingduplicate1() {
	a := []int{10, 20, 30, 30}
	b := set.TruncateList(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 4 [10 20 30 30]
}

func ExampleTruncateList_leadingduplicate() {
	a := []int{10, 10, 10, 20, 30}
	b := set.TruncateList(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 2 [20 30]
}

func ExampleTruncateList_leadingduplicate1() {
	a := []int{10, 10, 20, 30}
	b := set.TruncateList(a, 2)

	fmt.Printf(
		"%d %v",
		len(b),
		b,
	)
	// Output: 4 [10 10 20 30]
}

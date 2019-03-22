package common_test

import (
	"fmt"

	"github.com/EvanXzj/go-common"
)

// -----------------
// string.go
// -----------------

func ExampleReverse() {
	str := "!ooooooooG ,olleH"
	fmt.Println(common.Reverse(str))
	// Output: Hello, Goooooooo!
}

// -------END-------

// -----------------
// file.go
// -----------------

func ExampleIsFile() {
	fmt.Println(common.IsFile("file.go"))
	fmt.Println(common.IsFile("fake_path"))
	// Output:
	// true
	// false
}

// -------END-------

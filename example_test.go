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

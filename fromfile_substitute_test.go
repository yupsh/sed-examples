package sed_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/sed"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleSed_fromFile_substitute() {
	// sed 's/world/universe/' testdata/text.txt
	yup.MustRun(
		Sed(Pattern("s/world/universe/"), yup.File("testdata/text.txt")),
	)
	// Output:
	// hello universe
}


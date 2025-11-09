package sed_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sed"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleSed_fromFile_substitute() {
	// sed 's/world/universe/' testdata/text.txt
	gloo.MustRun(
		Sed(Pattern("s/world/universe/"), gloo.File("testdata/text.txt")),
	)
	// Output:
	// hello universe
}

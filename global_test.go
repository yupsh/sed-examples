package sed_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sed"
)

func ExampleSed_global() {
	// echo "foo foo foo" | sed 's/foo/bar/g'
	gloo.MustRun(
		Sed(Pattern("s/foo/bar/g"), strings.NewReader("foo foo foo")),
	)
	// Output:
	// bar bar bar
}

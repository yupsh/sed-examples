package sed_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/sed"
)

func ExampleSed_global() {
	// echo "foo foo foo" | sed 's/foo/bar/g'
	yup.MustRun(
		Sed(Pattern("s/foo/bar/g"), strings.NewReader("foo foo foo")),
	)
	// Output:
	// bar bar bar
}


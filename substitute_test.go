package sed_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/sed"
)

func ExampleSed_substitute() {
	// echo "hello world" | sed 's/world/universe/'
	yup.MustRun(
		Sed(Pattern("s/world/universe/"), strings.NewReader("hello world")),
	)
	// Output:
	// hello universe
}


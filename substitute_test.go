package sed_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/sed"
)

func ExampleSed_substitute() {
	// echo "hello world" | sed 's/world/universe/'
	gloo.MustRun(
		Sed(Pattern("s/world/universe/"), strings.NewReader("hello world")),
	)
	// Output:
	// hello universe
}

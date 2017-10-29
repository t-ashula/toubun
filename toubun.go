package main

import (
	"github.com/t-ashula/toubun/cmd"

	// for init()
	_ "github.com/t-ashula/toubun/runner"
	_ "github.com/t-ashula/toubun/runner/bundler"
	_ "github.com/t-ashula/toubun/runner/dep"
	_ "github.com/t-ashula/toubun/runner/gh"
	_ "github.com/t-ashula/toubun/runner/git"
	_ "github.com/t-ashula/toubun/runner/glide"
	_ "github.com/t-ashula/toubun/runner/npm"
)

func main() {
	cmd.Execute()
}

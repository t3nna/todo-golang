package main

import "github.com/hellflame/argparse"

type Opts struct {
	Add  []string
	List bool
	Comp []string
	Del  string
}

func getOpts() (Opts, error) {
	parser := argparse.NewParser("todo", "this is basic todo", nil)

	add := parser.Strings("a", "add", &argparse.Option{
		Required: false,
		Default:  "",
	})

	list := parser.Strings("l", "list", &argparse.Option{
		Required: false,
		Default:  "",
	})
	complete := parser.Strings("c", "complete", &argparse.Option{
		Required: false,
		Default:  "",
	})
	del := parser.Strings("d", "delete", &argparse.Option{
		Required: false,
		Default:  "",
	})

}

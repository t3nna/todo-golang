package main

import "github.com/hellflame/argparse"

type Opts struct {
	list bool
	add  []string
	comp []string
	rm   []string
}

func getOpts() (opts *Opts, err error) {
	parser := argparse.NewParser("todo", "gets all the values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	add := parser.Strings("a", "add", &argparse.Option{
		Default:  "",
		Required: false,
	})

	list := parser.Flag("l", "list", &argparse.Option{
		Required: false,
	})

	comp := parser.Strings("c", "comp", &argparse.Option{
		Default:  "",
		Required: false,
	})
	del := parser.Strings("d", "del", &argparse.Option{
		Default:  "",
		Required: false,
	})

	err = parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &Opts{
		add:  *add,
		list: *list,
		comp: *comp,
		rm:   *del,
	}, nil
}

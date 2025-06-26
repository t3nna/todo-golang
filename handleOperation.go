package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func handleOptions(opts *Opts, todo *Todo) error {
	if opts.list {
		handlePrinting(todo.data)
	} else if len(opts.add) > 0 {
		for _, v := range opts.add {
			newTodo := TodoItem{
				Description: v,
				CreatedAt:   time.Now().String(),
				IsComplete:  false,
			}
			todo.SetValue(&newTodo)
		}
		err := todo.Save()
		if err != nil {
			return err
		}
	} else if len(opts.rm) > 0 {
		for _, v := range opts.rm {
			todo.DeleteValue(v)
		}
	} else if len(opts.comp) > 0 {
		key := handlePrefixKey(todo.data, opts.comp)
		todo.MakeComplete(key)
	}

	defer handlePrinting(todo.data)

	return nil
}

func handlePrinting(data *Data) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintf(w, "ID\tTask\tCreated\tDone\n")
	for _, value := range data.Todo {

		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", value.ID, value.Description, value.CreatedAt, value.IsComplete)
	}
	w.Flush()
}

func handlePrefixKey(data *Data, key string) string {
	_, ok := data.Todo[key]
	if ok {
		return key
	}

	if len(key) < 4 {
		return ""
	}

	for k, _ := range data.Todo {
		if strings.HasPrefix(k, key) {
			return k
		}
	}
	return ""
}

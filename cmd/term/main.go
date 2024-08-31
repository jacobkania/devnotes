package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/jacobkania/devnotes/cmd/term/actions"
	"github.com/jacobkania/devnotes/config"
	"github.com/jacobkania/devnotes/db"
	"github.com/jacobkania/devnotes/migrate"
)

func main() {
	var flagToday bool

	flag.BoolVar(&flagToday, "t", false, "print all notes from today")

	flag.Parse()
	in := strings.Join(flag.Args(), " ")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Critical error: %v\n", r)
		}
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	d := db.Init(cfg, migrate.MigrateFS)

	if flagToday {
		fmt.Println("Today's notes:")

		err := actions.TodayNotes(d)
		if err != nil {
			fmt.Printf("DevNotes: Could not retrieve notes:\n%s\n", err)
		}
		return
	}

	if len(in) > 0 {
		err = actions.QuickNote(d, in)
		if err != nil {
			fmt.Printf("DevNotes: Could not create note:\n%s\n", err)
			return
		}
	}
}

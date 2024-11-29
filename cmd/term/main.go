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
	var flagStartWork bool
	var flagEndWork bool

	flag.BoolVar(&flagToday, "t", false, "print all notes from today")
	flag.BoolVar(&flagStartWork, "s", false, "start tracking working hours")
	flag.BoolVar(&flagEndWork, "e", false, "end tracking working hours")

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

	if flagToday || in == "today" {
		fmt.Println("Today's notes:")

		err := actions.TodayNotes(d)
		if err != nil {
			fmt.Printf("DevNotes Error: Could not retrieve notes: %s\n", err)
		}
	} else if flagStartWork {
		err = actions.WorkTrackingStart(d)
		if err != nil {
			fmt.Printf("DevNotes Error: Could not start work tracking: %s\n", err)
		}
	} else if flagEndWork {
		err = actions.WorkTrackingEnd(d)
		if err != nil {
			fmt.Printf("DevNotes Error: Could not end work tracking: %s\n", err)
		}
		
	} else if len(in) > 0 {
		err = actions.QuickNote(d, in)
		if err != nil {
			fmt.Printf("DevNotes Error: Could not create note: %s\n", err)
		}
	}
}

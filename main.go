package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron"
)

func main() {
	spec := flag.String("spec", "* * * * *", "cron spec")
	count := flag.Int("c", 5, "count")
	flag.Parse()

	if err := run(*spec, *count); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		os.Exit(1)
	}
}

func run(spec string, count int) error {
	s, err := cron.Parse(spec)
	if err != nil {
		return err
	}
	next := time.Now()
	for i := 0; i < count; i++ {
		next = s.Next(next)
		fmt.Println(next.Format(time.RFC3339))
	}
	return nil
}

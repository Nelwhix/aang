package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"github.com/joho/godotenv"
	"log"
)

func run(proj, commitMsg, stag string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	if proj == "" {
		return fmt.Errorf("staging directory is required: %w", ErrValidation)
	}

	if commitMsg == "" {
		return fmt.Errorf("commit message is required: %w", ErrValidation)
	}

	pipeline := make([]step, 4)

	pipeline[0] = newStep(
		"git add",
		"git",
		"Git add: SUCCESS",
		proj,
		"",
		[]string{"add", "."},
	)

	pipeline[1] = newStep(
		"git commit",
		"git",
		"Git commit: SUCCESS",
		proj,
		"",
		[]string{"commit", "-m", commitMsg},
	)

	pipeline[2] = newStep(
		"Generating static files",
		"npm",
		"Generating static files: SUCCESS",
		proj,
		"",
		[]string{"run", "generate"},
	)

	pipeline[3] = newStep(
		"Pushing to the Dev repo",
		"git",
		"Git push : SUCCESS",
		proj,
		stag,
		[]string{"push", "-u", "origin", "master"},
	)

	
	for _, s := range pipeline {
		msg, err := s.execute()
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err 
		}
	}
	
	return nil
}

func main() {
	proj := flag.String("p", "", "Project directory")
	msg := flag.String("m", "", "Commit message")
	stag := flag.String("s", "", "Staging directory")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	if err := run(*proj, *msg, *stag, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
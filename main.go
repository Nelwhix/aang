package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func run(proj, commitMsg string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	if commitMsg == "" {
		return fmt.Errorf("commit message is required: %w", ErrValidation)
	}

	f, err := os.Open(filepath.Join(proj, ".aang"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var pipeline []step
	for scanner.Scan() {
		command := scanner.Text()
		cmdArray := strings.Split(command, " ")
		exe := cmdArray[0]
		pipeline = append(pipeline, newStep(
			command,
			exe,
			command[0:2] + ": SUCCESS",
			proj,
			cmdArray[1:],
		),
		)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
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
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
		"aang-cli v0.0.1-alpha \n")
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright " + strconv.Itoa(time.Now().Local().Year()) + "\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
	proj := flag.String("p", "", "Project directory")
	msg := flag.String("m", "", "Commit message")
	flag.Parse()


	if err := run(*proj, *msg, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
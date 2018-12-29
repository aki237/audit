package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	help := flag.Bool("h", false, "Print help and exit")
	c := flag.String("c", "", "Path to the configuration file")
	flag.Parse()

	args := flag.Args()

	if *help || len(args) < 1 {
		flag.Usage()
		return
	}

	if *c == "" {
		fmt.Println("Running in default configuration")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	start := time.Now()
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Runtime : ", time.Since(start))

}

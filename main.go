package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "log", "--all -M -C --numstat --date=short --pretty=format:'--%h--%cd--%cn'")
	logs, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer logs.Close()
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	parser := NewParser(logs)
	entries, err := parser.Parse()
	fmt.Printf("%v", entries)
}

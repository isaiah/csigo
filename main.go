package main

import (
	"log"
	"os/exec"
)

func main() {
	git, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("installing git is in your future")
	}
	cmd := exec.Command(git, "log", "--all", "-M", "-C", "--numstat", "--date=short", "--pretty=format:'--%h--%cd--%cn'")
	logs, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	parser := NewParser(logs)
	entries, err := parser.Parse()
	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", entries)
}

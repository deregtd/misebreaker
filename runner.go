package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/alitto/pond/v2"
)

const numDirs = 2

func main() {
	pool := pond.NewPool(numDirs)
	group := pool.NewGroup()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	sp := path.Join(home, ".local", "share", "mise", "shims")
	os.Setenv("PATH", sp+":"+os.Getenv("PATH"))

	cmd := exec.Command("node", "--version")
	if err := cmd.Run(); err == nil {
		log.Fatal("remove node from mise global before running test")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Clean/setup
	if err := exec.Command("mise", "uninstall", "--all", "node").Run(); err != nil {
		log.Fatal(err)
	}

	for i := range numDirs {
		d := path.Join(wd, fmt.Sprintf("%d", i+1))
		cmd := exec.Command("mise", "trust")
		cmd.Dir = d
		cmd.Run()
	}

	for i := range numDirs {
		d := path.Join(wd, fmt.Sprintf("%d", i+1))
		group.Submit(func() {
			cmd := exec.Command("mise", "install")
			cmd.Dir = d
			op, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Got error on mise install #%d\n", i+1)
				fmt.Println(string(op))
				log.Fatal(err)
			}

			// fmt.Printf("PATH: %s\n", os.Getenv("PATH"))

			cmd = exec.Command("node", "--version")
			cmd.Dir = d
			op, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Got error on node version #%d\n", i+1)
				fmt.Println(string(op))
				log.Fatal(err)
			}
			fmt.Println("node: " + string(op))
		})
	}

	group.Wait()
}

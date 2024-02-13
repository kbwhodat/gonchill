package util

import (
	"os"
	"os/exec"
	"fmt"
	"log"
)

func Watch(selected_magnet string) {

	dir, err := exec.LookPath("peerflix")
	if err != nil {
		fmt.Println("Unable to find executable")
	}

	result := exec.Command(dir, selected_magnet, "-a", "-k")

	result.Stdin = os.Stdin
	result.Stdout = os.Stdout
	result.Stderr = os.Stderr

	if err := result.Run(); err != nil {
		log.Fatalf("Looks like something went wrong when trying to run Peeflix")
	}
}

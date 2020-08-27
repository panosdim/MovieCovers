package main

import (
	"log"
	"os"
	"os/exec"
)

func mp4Info(file string) string {
	out, err := exec.Command(atomicParsley, file, "-t").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func mp4SaveCover(file string, cover *os.File) {
	err := exec.Command(atomicParsley, file, "--artwork", cover.Name(), "--overWrite").Run()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"
	"os/exec"
)

func mp4Info(file string) string {
	// TODO: Check with LookPath if AtomicParsley is installed https://golang.org/pkg/os/exec/#example_LookPath
	out, err := exec.Command(`C:\Users\padi\Downloads\AtomicParsley-win32-0.9.0\AtomicParsley.exe`, file, "-t").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func mp4SaveCover(file string, cover *os.File) {
	err := exec.Command(`C:\Users\padi\Downloads\AtomicParsley-win32-0.9.0\AtomicParsley.exe`, file, "--artwork", cover.Name(), "--overWrite").Run()
	if err != nil {
		log.Fatal(err)
	}
}

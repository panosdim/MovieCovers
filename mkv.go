package main

import (
	"log"
	"os"
	"os/exec"
)

func mkvInfo(file string) string {
	out, err := exec.Command(mkvMerge, "-i", file).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func mkvSaveCover(file string, cover *os.File) {
	err := exec.Command(mkvPropEdit, file, "--attachment-name", "cover.jpg",
		"--attachment-mime-type", "image/jpeg",
		"--add-attachment", cover.Name()).Run()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"
	"os/exec"
)

func mkvInfo(file string) string {
	// TODO: Check with LookPath if mkvmerge is installed https://golang.org/pkg/os/exec/#example_LookPath
	out, err := exec.Command(`C:\Users\padi\Downloads\mkvtoolnix\mkvmerge.exe`, "-i", file).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func mkvSaveCover(file string, cover *os.File) {
	err := exec.Command(`C:\Users\padi\Downloads\mkvtoolnix\mkvpropedit.exe`, file, "--attachment-name", "cover.jpg",
		"--attachment-mime-type", "image/jpeg",
		"--add-attachment", cover.Name()).Run()
	if err != nil {
		log.Fatal(err)
	}
}

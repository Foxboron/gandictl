package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func TempFile(buf []byte) ([]byte, error) {
	tmpfile, err := ioutil.TempFile("/var/tmp", "goctl-*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(buf); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	if err := OpenFileInEditor(tmpfile.Name()); err != nil {
		return []byte{}, err
	}
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

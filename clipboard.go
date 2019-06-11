package main

import "os/exec"

func ReadClipboard() (string, error) {
	cmd := exec.Command("xclip", "-o")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

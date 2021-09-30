package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var out1, _ = exec.Command("ls").Output()
	fmt.Printf("-> ls :\n%s\n", string(out1))

	var out2, _ = exec.Command("pwd").Output()
	fmt.Printf("-> pwd :\n%s\n", string(out2))

	var out3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf("-> git config user.name :\n%s\n", string(out3))
	/*if runtime.GOOS == "windows" {
		output, err = exec.Command("CMD", "/C", "git config user.name").Output()
	} else {
		output, err = exec.Command("CMD", "-c", "git config user.name").Output()
	}*/
}

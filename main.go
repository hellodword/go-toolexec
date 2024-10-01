package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Fprintln(os.Stderr, "====> mytool", os.Args[1:])
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

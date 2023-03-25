package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        fmt.Println("Error:", lookErr)
    }

    args := []string{"ls", "-a", "-l", "-h"}

    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
			fmt.Println("Error:", execErr)
    }
}
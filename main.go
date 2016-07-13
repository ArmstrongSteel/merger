package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// runningApps stores the list of running apps to kill at the end
var runningApps []*exec.Cmd

func main() {
	// When the application is killed, make sure the child processes are killed too
	defer cleanup()

	// Make sure at least one app is sent as argument to run
	if len(os.Args) < 2 {
		fmt.Println("Please provide commands to run.")
	}

	// Loop through the name of apps to run and run them
	for index, arg := range os.Args {
		if index != 0 {
			// Split the commands by a empty space so that we can pass the arguments as slices to the execCommand function
			execCommand(strings.Split(arg, " ")...)
		}
	}

	// Run forever
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

// Kill all the child processes
func cleanup() {
	for _, cmd := range runningApps {
		cmd.Process.Kill()
	}
}

// Execute the command and output the stdout and stderr to the os stdout and stderr
func execCommand(commands ...string) {
	// Extract the first command as name and rest as commands (aka arguments)
	name, commands := commands[0], commands[1:]
	// Run commands
	cmd := exec.Command(name, commands...)

	// Pipe commands to terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Start command
	cmd.Start()

	// Add the programs to the list of running apps to kill at the end
	runningApps = append(runningApps, cmd)
}

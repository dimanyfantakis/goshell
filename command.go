package main

import (
	"os"
	"os/exec"
	"strings"
)

type Command interface {
	Execute(args []string)
}

type BuiltinCommand struct {
	Exec func(args []string, wd *string)
}

func (c BuiltinCommand) Execute(args []string, wd *string) {
	c.Exec(args, wd)
}

var builtInCommands map[string]BuiltinCommand

func cd(args []string, wd *string) {
	if len(args) != 2 {
		panic("cd takes exactly 1 argument")
	}
	os.Chdir(args[1])
	newDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	*wd = newDirectory
}

func exit(args []string, wd *string) {
	if len(args) != 1 {
		panic("exit doesn't take any arguments")
	}
	os.Exit(0)
}

func init() {
	builtInCommands = map[string]BuiltinCommand{
		"cd": {
			Exec: cd,
		},
		"exit": {
			Exec: exit,
		},
	}
}

func Execute(input string, wd *string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	builtInCommand, exists := builtInCommands[args[0]]
	if exists {
		builtInCommand.Execute(args, wd)
		return nil
	}

	command := exec.Command(args[0], args[1:]...)

	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	return err
}

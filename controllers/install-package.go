package controllers

import (
	"ghpm/models"
	"os"
	"os/exec"
	"strings"
)

func InstallPackage(pkg models.Package) {
	var commands [][]string
	var i int = 0
	for i < len(pkg.Instructions) {
		command := strings.Split(pkg.Instructions[i], " ")
		commands = append(commands, command)
		i++
	}
	i = 0
	for i < len(commands) {
		a := commands[i][0]
		commands[i] = commands[i][1:]
		cmd := exec.Command(a, commands[i]...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
		i++
	}
}

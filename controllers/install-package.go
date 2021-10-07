package controllers

import (
	"fmt"
	"ghpm/config"
	"ghpm/models"
	"os"
	"os/exec"
	"strings"
)

func InstallPackage(url string, list models.LicensedPackages) {
	isAuth := CheckAuth(url, list)
	if !isAuth {
		fmt.Println("This package isn't licensed. Are you sure you want to install it?")
		fmt.Println("Yes[Y]/No[N]:")
		var answer string
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)
		if answer == "N" || answer == "n" {
			return
		}
	}
	cmd := exec.Command("git", "clone", url)
	nameOfdir := strings.Split(url, "/")[4]
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Chdir(nameOfdir)
	fmt.Println(nameOfdir)
	pkg := config.LoadPackageInfo()
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

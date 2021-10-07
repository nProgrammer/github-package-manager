/*
	Norbert Wagner (nProgrammer) 2021
	GHPM - GitHub Package Manager
*/

package main

import (
	"flag"
	"fmt"
	"ghpm/config"
	"ghpm/controllers"
)

func main() {
	authP := config.LoadAuthorizedPackagesList()
	check := flag.Bool("check-auth", false, "Function that is checking is package authorized in official database")
	ins := flag.Bool("install", false, "Function that is installing package")
	help := flag.Bool("help", false, "Showing help section")
	showAuth := flag.Bool("auth-packages", false, "Showing authorized packages")
	url := flag.String("url", "", "Url of package")
	flag.Parse()

	if *check {
		answ := controllers.CheckAuth(*url, authP)
		if answ {
			fmt.Println("Package is safe")
		} else {
			fmt.Println("Package isn't register in database")
		}
	}
	if *showAuth {
		controllers.ShowAuthorizedPackages(authP)
	}
	if *ins {
		controllers.InstallPackage(*url, authP)
	}

	if *help {
		showHelp()
	}
}

func showHelp() {
	flag.Usage()
}

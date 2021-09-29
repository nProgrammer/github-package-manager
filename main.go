/*
	Norbert Wagner (nProgrammer) 2021
	GHPM - GitHub Package Manager
*/

package main

import (
	"flag"
	"ghpm/config"
	"ghpm/controllers"
)

func main() {
	authP := config.LoadAuthorizedPackagesList()
	check := flag.Bool("check-auth", false, "Function that is checking is package authorized in official database")
	help := flag.Bool("help", false, "Showing help section")
	url := flag.String("url", "", "Url of package")
	flag.Parse()

	if *check {
		controllers.CheckAuth(*url, authP)
	}

	if *help {
		showHelp()
	}
}

func showHelp() {
	flag.Usage()
}

package controllers

import (
	"fmt"
	"ghpm/models"
)

func ShowAuthorizedPackages(list models.LicensedPackages) {
	fmt.Println("List of authorized packages:")
	i := 0
	for i < len(list.Packages) {
		fmt.Println("- ", list)
		i++
	}
}

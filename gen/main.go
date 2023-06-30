package main

import (
	"fmt"
	"gen/lang"
	"gen/models"
	"os"
)

func main() {
	all := len(os.Args) < 2 || os.Args[1] == "all"
	used := false
	if all || os.Args[1] == "lang" {
		fmt.Println("Generating language file")
		lang.Generate()
		used = true
	}
	if all || os.Args[1] == "models" {
		fmt.Println("Generating model files")
		models.Generate()
		used = true
	}
	if !used {
		fmt.Println("Please pass one of the commands to the program:")
		fmt.Println("\t   all - generate both lang and model files")
		fmt.Println("\t  lang - generate only language file")
		fmt.Println("\tmodels - generate only models file")
		os.Exit(1)
	}
}

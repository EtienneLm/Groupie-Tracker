package groupietrackers

import (
	"bufio"
	"fmt"
	"os"
)

var SpotifyId []string

func ArrayInit(DB string) []string { //This function take a DB name and return an array from the DB
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	file, err := os.Open(DB)
	if err != nil { //We check if the data base is correct
		fmt.Println(colorRed, "ERROR no data base named : ", DB, colorReset)
		os.Exit(1) //We stop the programme
	}
	var array []string //We create the array who will contain the entire words data base

	scanner := bufio.NewScanner(file) //We take make a array with all the lines of the DB

	for scanner.Scan() { //We explore the scanner
		if scanner.Text() != "" || DB == "DB/DBSave.txt" {
			array = append(array, scanner.Text()) //We append the DB line by line in the array variable
		}
	}
	return array
}

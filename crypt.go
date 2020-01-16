package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var flagPassword = flag.String("s", "", "string to bcrypt")
var cryptStrength = flag.Int("cost", 12, "bcrypt cost value")

/**
* Handy bCrypt tool for hashing strings with bcrypt by command line
*
 */
func main() {
	flag.Parse()

	if *flagPassword == "" {
		fmt.Printf("[crypt] error no input string supplied (-s)")
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(*flagPassword), *cryptStrength)
	if err != nil {
		fmt.Printf("[crypt] failed to bcrypt string :%v\n", err)
	}
	fmt.Fprintf(os.Stdout, "%s", bytes)

	return
}

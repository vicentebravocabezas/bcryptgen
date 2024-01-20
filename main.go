package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	dFlag := false
	flag.BoolVar(&dFlag, "d", false, "Compare a hash")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage:\nbcrypt <password>\nbcrypt -d <password> '<hash>'")
		return
	}

	if dFlag && len(os.Args) < 4 {
		fmt.Println("Specify a password followed by a hash")
		return
	}

	if dFlag {
		fmt.Println(os.Args[3])
		if err := bcrypt.CompareHashAndPassword([]byte(os.Args[3]), []byte(os.Args[2])); err != nil {
			fmt.Println("Password comparison failed")
			return
		}
		fmt.Println("Comparison successful")
		return
	}

	password := os.Args[1]

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hash))
}
